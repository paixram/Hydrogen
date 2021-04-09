package object

import "fmt"

type Environment struct {
	store  map[string]Object
	blocks map[string]Object
	outer  *Environment
}

func NewEnvironemnt() *Environment {
	s := make(map[string]Object)
	b := make(map[string]Object)
	return &Environment{store: s, blocks: b}
}

func (e *Environment) Get(typeEnv string, name string) (Object, bool) {
	var obj Object
	var ok bool
	if typeEnv == "store" {
		obj, ok = e.store[name]

		//fmt.Println("El pepe: ", obj.(*Macros).Value, ok)
		if !ok && e.outer != nil {
			obj, ok = e.outer.Get("store", name)
		}
	} else if typeEnv == "block" {
		obj, ok = e.blocks[name]
		if !ok && e.outer != nil {
			obj, ok = e.outer.Get("block", name)
		}
	}

	return obj, ok
}

func (e *Environment) Set(typeEnv string, name string, val Object) Object {
	if typeEnv == "store" {
		if obj, ok := e.store[name]; ok {
			if obj.Type() == MACROS_OBJ {
				//log.Fatal("xdd macos")
				return &Error{Message: fmt.Sprintf("You cannot redeclare or use the name of the macro: %s", name)}
				//return err.NewError("Is a macros dont cvhasnger vallue"), false
			}
		}
		e.store[name] = val
	} else if typeEnv == "block" {
		e.blocks[name] = val
	}

	return val
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironemnt()
	env.outer = outer
	return env
}
