package object

import (
	"bytes"
	"strings"

	"github.com/Hydrogen/src/ast"
	"github.com/Hydrogen/src/token"
)

type Function struct {
	Parameters []*ast.Identifier
	Name       string
	Body       *ast.BlockStatement
	Env        *Environment
	ReturnType token.TokenType
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}
