package ast

import (
	"bytes"

	"github.com/Hydrogen/src/token"
)

type BlockLiteral struct {
	Token token.Token
	Name  *Identifier
	Body  *BlockStatement
}

func (b *BlockLiteral) expressionNode()      {}
func (b *BlockLiteral) TokenLiteral() string { return b.Token.Literal }
func (b *BlockLiteral) String() string {
	var out bytes.Buffer

	for _, tt := range b.Body.Statements {
		out.WriteString(tt.String())
		out.WriteString("sdssdf")
	}
	out.WriteString("sdssdf")
	return out.String()
}
