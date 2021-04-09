package ast

import "github.com/Hydrogen/src/token"

type TypeLiteral struct {
	Token token.Token
	Value string
}

func (s *TypeLiteral) expressionNode()      {}
func (s *TypeLiteral) TokenLiteral() string { return s.Token.Literal }
func (s *TypeLiteral) String() string {
	return "TypeLiteral"
}
