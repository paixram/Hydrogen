package ast

import (
	"bytes"

	"github.com/Hydrogen/src/token"
)

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (r *ReturnStatement) statementNode()       {}
func (r *ReturnStatement) TokenLiteral() string { return r.Token.Literal }
func (r *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(r.TokenLiteral() + " ")
	out.WriteString(r.Value.String() + ";")

	return out.String()
}
