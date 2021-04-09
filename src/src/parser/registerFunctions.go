package parser

import (
	"github.com/Hydrogen/src/token"
)

func (p *Parser) registerPrefix(t token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[t] = fn
}

func (p *Parser) infixRegister(t token.TokenType, fn infixParseFn) {
	p.infixParseFns[t] = fn
}
