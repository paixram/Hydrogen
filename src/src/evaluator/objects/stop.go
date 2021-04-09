package object

type Stop struct {
	Value Object
}

func (rv *Stop) Type() ObjectType { return STOP_OBJ }
func (rv *Stop) Inspect() string  { return "Stop" }
