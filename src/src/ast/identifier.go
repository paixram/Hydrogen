package ast

import "github.com/Hydrogen/src/token"

type Identifier struct {
	Token     token.Token
	Value     string
	TypeValue token.TokenType
}

func (i *Identifier) String() string { return i.Value }

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
