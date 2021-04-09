package object

type Macros struct {
	Name  string
	Value Object
}

func (m *Macros) Type() ObjectType { return MACROS_OBJ }
func (m *Macros) Inspect() string  { return m.Value.Inspect() }
