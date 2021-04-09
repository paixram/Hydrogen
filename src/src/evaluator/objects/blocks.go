package object

import (
	"bytes"

	"github.com/Hydrogen/src/ast"
)

type Block struct {
	Body *ast.BlockStatement
	Env  *Environment
}

func (f *Block) Type() ObjectType { return FUNCTION_OBJ }
func (f *Block) Inspect() string {
	var out bytes.Buffer

	return out.String()
}
