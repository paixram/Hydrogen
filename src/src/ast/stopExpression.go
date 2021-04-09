package ast

import "github.com/Hydrogen/src/token"

type StopExpression struct {
	Token token.Token
}

func (s *StopExpression) expressionNode()      {}
func (s *StopExpression) TokenLiteral() string { return s.Token.Literal }
func (s *StopExpression) String() string {
	return s.Token.Literal
}
