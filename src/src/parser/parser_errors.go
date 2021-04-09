package parser

import (
	"fmt"

	"github.com/Hydrogen/src/token"
)

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("Expected next token to be %s, got %s INSTEAD",
		t, p.peekToken.Type)

	p.errors = append(p.errors, msg)
}

func (p *Parser) noPrefixParserFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)

}
