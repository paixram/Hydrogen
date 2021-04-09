package ast

import (
	"bytes"

	"github.com/Hydrogen/src/token"
)

type DeclareStatement struct {
	Token token.Token
	Name  *Identifier
	Type  *TypeLiteral
	Value Expression
}

func (d *DeclareStatement) statementNode()       {}
func (d *DeclareStatement) TokenLiteral() string { return d.Token.Literal }
func (d *DeclareStatement) String() string {
	var out bytes.Buffer

	out.WriteString(d.TokenLiteral())
	out.WriteString(" " + d.Name.Value + " ")
	out.WriteString(string(d.Type.Token.Type) + " ")
	out.WriteString(token.ASSIGN + " ")
	out.WriteString(d.Value.String() + ";")

	return out.String()
}
