package object

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

type BuiltinFunction func(args ...Object) Object
