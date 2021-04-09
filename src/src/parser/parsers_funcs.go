package parser

import (
	"fmt"
	"strconv"

	"github.com/Hydrogen/src/ast"
	"github.com/Hydrogen/src/token"
)

func (p *Parser) parseDeclareStatement() *ast.DeclareStatement {
	dec := &ast.DeclareStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		p.peekError(token.IDENT)
		return nil
	}
	dec.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.INT) && !p.expectPeek(token.STRING) && !p.expectPeek(token.FLOAT) && !p.expectPeek(token.BOOL) {
		p.peekError(token.TYPE)
		return nil
	}
	dec.Name.TypeValue = p.curToken.Type

	dec.Type = &ast.TypeLiteral{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		p.peekError(token.ASSIGN)
		return nil
	}
	p.nextToken()

	// TODO: Crear el valor de la variable y guardarlo
	if dec.Type.Value == "bool" {
		if !p.curTokenIs(token.TRUE) && !p.curTokenIs(token.FALSE) {
			msg := fmt.Sprintf("Incorrect Boolean values ​​for the variable: %s is %s data type and get %s value",
				dec.Name.Value, dec.Name.TypeValue, p.curToken.Type)
			p.errors = append(p.errors, msg)
		}
	} else if dec.Name.TypeValue != p.curToken.Type {
		msg := fmt.Sprintf("Uneven data types for the variable: %s, the data type of the variable is %s and it was given a type %s Value: %s",
			dec.Name.Value, dec.Name.TypeValue, p.curToken.Type, p.curToken.Literal)
		p.errors = append(p.errors, msg)
	}

	dec.Value = p.parseExpression(LOWEST)

	for p.peekTokenIs(token.SEMICOLON) {
		//msg := fmt.Sprintf("Ejecutando ando")
		//p.errors = append(p.errors, msg)
		p.nextToken()
	}
	//msg := fmt.Sprintf("Ejecutando ando")
	//p.errors = append(p.errors, msg)
	return dec
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	rtrStmt := &ast.ReturnStatement{Token: p.curToken}

	if !p.expectPeek(token.INT) && !p.expectPeek(token.STRING) && !p.expectPeek(token.FLOAT) && !p.expectPeek(token.BOOL) && !p.expectPeek(token.IDENT) {
		rtrStmt.Value = nil
		return rtrStmt
	}

	// TODO: crear el valor de retorno para el return statement
	rtrStmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return rtrStmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParserFnError(p.curToken.Type)
		return nil
	}

	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

// PREFIX AND INFIX FUNCTION
func (p *Parser) parseBlock() ast.Expression {
	block := &ast.BlockLiteral{Token: p.curToken}
	block.Name = &ast.Identifier{TypeValue: p.curToken.Type}
	if !p.expectPeek(token.IDENT) {
		p.peekError(token.IDENT)
		return nil
	}

	block.Name.Token = p.curToken
	block.Name.Value = p.curToken.Literal
	//block.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.LBRACE) {
		p.peekError(token.LBRACE)
		return nil
	}

	block.Body = p.parseBlockStatement()

	return block
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	bodyBlock := &ast.BlockStatement{Token: p.curToken}
	bodyBlock.Statements = []ast.Statement{}

	p.nextToken()

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		if p.curTokenIs(token.RETURN) {
			msg := fmt.Sprintf("Can't use RETURN reserved word in Block")
			// TODO: CREAR UN LOG FATAL PARA INDICAR QUE NO SE PUEDE PONER RETURN EN UN BLOQUE
			p.errors = append(p.errors, msg)
		}

		stmt := p.parseStatement()
		if stmt != nil {
			bodyBlock.Statements = append(bodyBlock.Statements, stmt)
		}
		p.nextToken()
	}

	return bodyBlock
}

func (p *Parser) parseIdentificator() ast.Expression {
	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal, TypeValue: token.IDENT}
	//fmt.Println("IDENT: ", ident.Value)
	return ident
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	intLit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	intLit.Value = value

	return intLit
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.BooleanLiteral{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	//defer untrace(trace("parsePrefixExpression"))

	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *Parser) parseCallBlock(fu ast.Expression) ast.Expression {
	fu.(*ast.Identifier).TypeValue = token.BLOCK
	callBlock := &ast.CallBlockExpression{Token: p.curToken, Block: fu}
	//fmt.Println("Calñlblock:", *callBlock.Block.(*ast.Identifier))
	return callBlock // TODO: llamar a caal blockm
}

func (p *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		msg := fmt.Sprintf("This IF statement is not valid")
		p.errors = append(p.errors, msg)
		return nil
	}

	p.nextToken()
	expression.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		msg := fmt.Sprintf("Missing close the expression if with the sign ')'")
		p.errors = append(p.errors, msg)
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	expression.Consequence = p.parseBlockStatementSortinf()

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()

		if !p.expectPeek(token.LBRACE) {
			return nil
		}

		expression.Alternative = p.parseBlockStatementSortinf()
	}

	return expression
}

func (p *Parser) parseStop() ast.Expression {
	return &ast.StopExpression{Token: p.curToken}
}

func (p *Parser) parseBlockStatementSortinf() *ast.BlockStatement {
	funBod := &ast.BlockStatement{Token: p.curToken}

	funBod.Statements = []ast.Statement{}

	p.nextToken()

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			funBod.Statements = append(funBod.Statements, stmt)
		}
		p.nextToken()
	}

	return funBod
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
	funquer := &ast.FunctionLiteral{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		msg := fmt.Sprintf("The function:% s does not have an identified name",
			p.peekToken.Literal)
		p.errors = append(p.errors, msg)
		// TODO: retornar nil
	}
	// TODO GET NAME
	funquer.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.LPAREN) {
		msg := fmt.Sprintf("Function syntax error:% s",
			p.curToken.Literal)
		p.errors = append(p.errors, msg)
		// TODO: retornar nil
	}

	funquer.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(token.DEVOLUCION) {
		msg := fmt.Sprintf("The function not have type return")
		p.errors = append(p.errors, msg)
	}
	p.nextToken()
	// TODO: hacer tipos dinamicos en el futuro
	if !p.curTokenIs(token.INT) && !p.curTokenIs(token.STRING) && !p.curTokenIs(token.FLOAT) && !p.curTokenIs(token.BOOL) {
		msg := fmt.Sprintf("Is not valid tipe")
		p.errors = append(p.errors, msg)
		return nil
	}
	funquer.Name.TypeValue = p.curToken.Type
	//fmt.Println(p.curToken)
	if !p.expectPeek(token.LBRACE) {
		// TODO: HANLDE LBRACE ERROR
		return nil
	}

	funquer.Body = p.parseBlockStatementSortinf()

	return funquer
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	// TODO: verificar si existe un tipo de dato apra la variable
	p.nextToken()
	ident.TypeValue = p.curToken.Type
	identifiers = append(identifiers, ident)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		p.nextToken()
		ident.TypeValue = p.curToken.Type
		identifiers = append(identifiers, ident)
	}

	if !p.expectPeek(token.RPAREN) {
		msg := fmt.Sprintf("The Sign does not RPAREN")
		p.errors = append(p.errors, msg)
		return nil
	}

	return identifiers
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	function.(*ast.Identifier).TypeValue = token.FUNCTION
	exp := &ast.CallExpression{Token: p.curToken, Function: function}
	exp.Arguments = p.parseCallArguments()
	return exp
}

func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	return args
}

func (p *Parser) parseMacrosDef() *ast.Macros {
	def := &ast.Macros{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		msg := fmt.Sprint("A macros recive IDENT after DEF keyword")
		p.errors = append(p.errors, msg)
		return nil
	}
	//fmt.Println(def, p.curToken)
	def.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal, TypeValue: token.DEF}

	p.nextToken() // Temporal hasta usar expect peek para coimprobar tipos
	// TODO: COMPROBAR LOS TIPOD DE DATOS EN VALUE
	def.Value = p.parseExpression(LOWEST)
	//fmt.Println("Valor", def.Value)
	return def
}
