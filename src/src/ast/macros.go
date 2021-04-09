package ast

import (
	"github.com/Hydrogen/src/token"
)

type Macros struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (d *Macros) statementNode()       {}
func (d *Macros) TokenLiteral() string { return d.Token.Literal }
func (d *Macros) String() string {
	/*var out bytes.Buffer

	out.WriteString(d.TokenLiteral())
	out.WriteString(" " + d.Name.Value + " ")
	out.WriteString(string(d.String()) + " ")
	out.WriteString(token.ASSIGN + " ")
	out.WriteString(d.Value.String() + ";")

	return out.String()*/
	return "MACROS"
}
