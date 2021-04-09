package builts

import (
	"fmt"

	err "github.com/Hydrogen/src/errors"
	object "github.com/Hydrogen/src/evaluator/objects"
)

var Builtins = map[string]*object.Builtin{
	"println": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return err.NewError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.Integer:
				fmt.Println(arg.Value)
			case *object.String:
				fmt.Println(arg.Value)
			case *object.Boolean:
				fmt.Println(arg.Value)
			}
			return nil
		},
	},
}
