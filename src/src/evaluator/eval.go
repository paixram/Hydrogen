package evaluator

import (
	"github.com/Hydrogen/src/ast"
	object "github.com/Hydrogen/src/evaluator/objects"
)

func isError(obj object.Object) bool {
	//fmt.Println("Obreron: ", obj)
	if obj != nil {
		//fmt.Println("Obreron: ", obj)
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.BooleanLiteral:
		return nativeBooleanObject(node.Value)
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	case *ast.Identifier:
		s := evalIdentifier(node, env)
		if isError(s) {
			//fmt.Println(s)
			return s
		}
		//fmt.Println(s)
		return s
	case *ast.IfExpression:
		return evalIfExpression(node, env)
	case *ast.StopExpression:
		return &object.Stop{Value: &object.Boolean{Value: true}}
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)
	case *ast.BlockLiteral:
		body := node.Body

		env.Set("block", node.Name.Value, &object.Block{Body: body})
		//return &object.Block{Body: body}
	case *ast.CallBlockExpression:
		blocking := Eval(node.Block, env)
		//fmt.Println(blocking)
		return applyBlock(blocking, env)
	case *ast.DeclareStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}

		e := env.Set("store", node.Name.Value, val)
		if isError(e) {
			return e
		}
	case *ast.Macros:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		d := &object.Macros{Name: node.Name.Value, Value: val}
		e := env.Set("store", node.Name.Value, d)
		if isError(e) {
			return e
		}
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *ast.ReturnStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}
	case *ast.FunctionLiteral:
		name := node.Name.Value
		params := node.Parameters
		body := node.Body
		returnTyper := node.Name.TypeValue
		env.Set("store", node.Name.Value, &object.Function{Parameters: params, Body: body, ReturnType: returnTyper, Name: name})
	case *ast.CallExpression:
		//devolution := node.Function.(*ast.Identifier)
		//fmt.Println(devolution)
		funct := Eval(node.Function, env)
		//fmt.Println("Function", funct)
		if isError(funct) {
			return funct
		}

		args := evalExpressions(node.Arguments, env)

		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		return applyFunction(funct, args)
	}

	return nil
}
