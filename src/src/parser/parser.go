package parser

import (
	"github.com/Hydrogen/src/ast"
	"github.com/Hydrogen/src/lexer"
	"github.com/Hydrogen/src/token"
)

// fUNCS CORE
type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l *lexer.Lexer

	// ERROR ANALITICS
	errors []string

	curToken  token.Token
	peekToken token.Token

	// MPPING
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(input *lexer.Lexer) *Parser {
	p := &Parser{l: input, errors: []string{}}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentificator)
	p.registerPrefix(token.PLUS, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.STRING, p.parseStringLiteral)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.IF, p.parseIfExpression)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(token.BLOCK, p.parseBlock)
	p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral)

	p.registerPrefix(token.STOP, p.parseStop)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.infixRegister(token.PLUS, p.parseInfixExpression)
	p.infixRegister(token.MINUS, p.parseInfixExpression)
	p.infixRegister(token.SLASH, p.parseInfixExpression)
	p.infixRegister(token.ASTERISK, p.parseInfixExpression)
	p.infixRegister(token.EQ, p.parseInfixExpression)
	p.infixRegister(token.NOT_EQ, p.parseInfixExpression)
	p.infixRegister(token.LT, p.parseInfixExpression)
	p.infixRegister(token.GT, p.parseInfixExpression)
	p.infixRegister(token.LPAREN, p.parseCallExpression)
	p.infixRegister(token.INTERROG, p.parseCallBlock)

	// INICITIALIZE TOKENIZER
	p.nextToken()
	p.nextToken()

	return p
}

// CENTRAL CORE OF PARSER
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// TODO: hay que verificar que stop este bien
	// y el operador || por que puede que haiga errores
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program
}

// 2COND CORE
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.DEC:
		return p.parseDeclareStatement()
	case token.DEF:
		return p.parseMacrosDef()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
