package evaluator

import (
	"fmt"

	"github.com/Hydrogen/src/ast"
	builts "github.com/Hydrogen/src/builtins"
	err "github.com/Hydrogen/src/errors"
	object "github.com/Hydrogen/src/evaluator/objects"
	"github.com/Hydrogen/src/token"
)

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range program.Statements {
		result = Eval(statement, env)
		switch result := result.(type) {
		case *object.Error:
			return result
		case *object.Stop:
			return result.Value
		case *object.ReturnValue:
			return result.Value
		}
	}

	return result
}

func nativeBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}

	return FALSE
}

func evalIdentifier(
	node *ast.Identifier,
	env *object.Environment,
) object.Object {
	//fmt.Println(node.TypeValue)
	switch node.TypeValue {
	case token.BLOCK:
		//fmt.Println(node.TypeValue)
		if val, ok := env.Get("block", node.Value); ok {
			return val
		}
	case token.FUNCTION:
		//fmt.Println(node.TypeValue)
		if builtin, ok := builts.Builtins[node.Value]; ok {
			return builtin
		}

		if val, ok := env.Get("store", node.Value); ok {
			return val
		}
	case token.IDENT:
		if val, ok := env.Get("store", node.Value); ok {
			return val
		}
		//fmt.Println("Error", node.TypeValue)
	default:
		return err.NewError("identifier not found: " + node.Value)
	}
	return err.NewError("identifier not found: " + node.Value)
}

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {

	condition := Eval(ie.Condition, env)

	if isError(condition) {
		return condition
	}

	if isTrusthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return TRUE
	}
}

func isTrusthy(obj object.Object) bool {
	switch obj {
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range block.Statements {
		result = Eval(statement, env)
		if result != nil {
			rt := result.Type()
			// TODO: crear el STOP para que se detenga
			if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ || rt == object.STOP_OBJ {
				return result
			}
		}
	}

	return result
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	// TODO: optimizar esta parte
	if left.Type() == object.MACROS_OBJ {
		left = left.(*object.Macros).Value
	} else if right.Type() == object.MACROS_OBJ {
		right = right.(*object.Macros).Value
	}
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case operator == "==":
		return nativeBooleanObject(left == right)
	case operator == "!=":
		return nativeBooleanObject(left != right)
	case left.Type() != right.Type():
		return err.NewError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	default:
		return err.NewError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}

}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "<":
		return nativeBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBooleanObject(leftVal > rightVal)
	case "==":
		return nativeBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBooleanObject(leftVal != rightVal)
	default:
		return err.NewError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusOperatorExpression(right)
	default:
		return err.NewError("unknown operator %s%s", operator, right.Type())
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	default:
		// TODO: hacer rror
		return nil
	}
}

func evalMinusOperatorExpression(right object.Object) object.Object {

	if right.Type() != object.INTEGER_OBJ {
		return err.NewError("unknown operator: -%s", right.Type())
	}

	if right.Type() != object.INTEGER_OBJ {
		return nil
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func applyBlock(bl object.Object, env *object.Environment) object.Object {
	switch bl := bl.(type) {
	case *object.Block:
		extendedEnv := object.NewEnclosedEnvironment(env)
		evaluated := Eval(bl.Body, extendedEnv)
		return unwrapStop(evaluated)
	default:
		return err.NewError("not a function or block: %s", bl.Type())
	}
}

func unwrapStop(obj object.Object) object.Object {
	if _, ok := obj.(*object.Stop); ok {
		return nil
	}

	return obj
}

func evalExpressions(
	exps []ast.Expression,
	env *object.Environment,
) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		//fmt.Println("args", *&evaluated.(*object.Macros).Value)
		if evaluated.Type() == object.MACROS_OBJ {
			evaluated = evaluated.(*object.Macros).Value
		}

		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		extendedEnv, err := extendFunctionEnv(fn, args)
		if err != nil {
			return err
		}
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated, fn.ReturnType, fn.Name)
	case *object.Builtin:
		return fn.Fn(args...)
	default:
		return err.NewError("not a function: %s", fn.Type())
	}
}

func extendFunctionEnv(
	fn *object.Function,
	args []object.Object,
) (*object.Environment, object.Object) {
	env := object.NewEnclosedEnvironment(fn.Env)

	for paramIdx, param := range fn.Parameters {
		//fmt.Println(param.TypeValue, args[paramIdx].Type())
		if param.TypeValue != token.TokenType(args[paramIdx].Type()) {
			msg := fmt.Sprintf("The parameter (%s) is type %s and have %s", param.Token.Literal, param.TypeValue, args[paramIdx].Type())
			return nil, err.NewError(msg)
			//fmt.Println(msg)
			//os.Exit(1)
		}
		env.Set("store", param.Value, args[paramIdx])
	}

	return env, nil
}

func unwrapReturnValue(obj object.Object, returnType token.TokenType, fnName string) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		//fmt.Println(returnValue.Value.Type(), returnType)
		if returnType != token.TokenType(returnValue.Value.Type()) {
			msg := fmt.Sprintf("Return value of function %s is %s and returns %s, return values ​​are not of the same type",
				fnName, returnType, returnValue.Value.Type())
			return err.NewError(msg)
		}
		return returnValue.Value
	}
	return obj
}
