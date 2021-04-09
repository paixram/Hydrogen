package ast

import (
	"bytes"

	"github.com/Hydrogen/src/token"
)

type CallBlockExpression struct {
	Token token.Token
	Block Expression // IDENTIFIER
}

func (cb *CallBlockExpression) expressionNode()      {}
func (cb *CallBlockExpression) TokenLiteral() string { return cb.Token.Literal }
func (cb *CallBlockExpression) String() string {
	var out bytes.Buffer
	out.WriteString(cb.Block.String())
	return out.String()
}
