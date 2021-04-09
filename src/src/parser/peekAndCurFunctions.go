package parser

import "github.com/Hydrogen/src/token"

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) curPrecedence() int {
	if obj, ok := precedences[p.curToken.Type]; ok {
		return obj
	}

	return LOWEST
}

func (p *Parser) peekPrecedence() int {
	if obj, ok := precedences[p.peekToken.Type]; ok {
		return obj
	}

	return LOWEST
}
