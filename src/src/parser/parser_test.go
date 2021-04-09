package parser

import (
	"testing"

	"github.com/Hydrogen/src/ast"
	"github.com/Hydrogen/src/lexer"
	"github.com/Hydrogen/src/token"
)

func TestArtimetic(t *testing.T) {
	test := []struct {
		input    string
		expected string
	}{
		{`add(a + b + c * d / f + g)`, "add((((a + b) + ((c * d) / f)) + g))"},
	}

	for _, tt := range test {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()

		stmt := program.Statements[0].(*ast.ExpressionStatement)

		stmte := stmt.Expression.(*ast.CallExpression)

		if tt.expected != stmte.String() {
			t.Errorf("Expected=%s, Got=%s", tt.expected, stmte.String())
		}
	}
}

func TestCallExpression(t *testing.T) {
	test := []struct {
		input    string
		expected string
	}{
		{`add(x, t);`, "add()"},
		{`add(a + b + c * d / f + g)`, "add((((a + b) + ((c * d) / f)) + g))"},
	}

	for _, tt := range test {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		stmt := program.Statements[0].(*ast.ExpressionStatement)

		stmte := stmt.Expression.(*ast.CallExpression)

		if tt.expected != stmte.String() {
			t.Errorf("Expected=%s, Got=%s", tt.expected, stmte.String())
		}
	}
}

func TestFunction(t *testing.T) {
	test := []struct {
		input     string
		exprected string
	}{
		{`fn main() <- int {dec xd int -> 23;}`, "main"},
	}

	for _, tt := range test {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		stmt := program.Statements[0].(*ast.ExpressionStatement)

		stmte := stmt.Expression.(*ast.FunctionLiteral)
		if tt.exprected != stmte.Name.Value {
			t.Errorf("Expected=%s, Got=%s", tt.exprected, stmte.Name.Value)
		}
	}
}

func TestBlockExpression(t *testing.T) {
	test := []struct {
		input     string
		exprected string
	}{
		{`Block createVariable{dec integro int -> 23; STOP}`, "return 23;"},
	}

	for _, tt := range test {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		stmt := program.Statements[0].(*ast.ExpressionStatement)

		stmte := stmt.Expression.(*ast.BlockLiteral)

		if tt.exprected != stmte.String() {
			t.Errorf("Expected=%s, Got=%s", tt.exprected, stmte.String())
		}
	}
}

func TestIfExpression(t *testing.T) {
	test := []struct {
		input   string
		exprect string
	}{
		{`if(3 < 5){return 34;}`, "if(3 < 5) return 34;"},
		{`if(4 < 6){Block corefile{STOP;}; ?corefile}`, "if(4 < 6) Blockcorefile"},
	}

	for _, tt := range test {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		stmt := program.Statements[0].(*ast.ExpressionStatement)

		stmte := stmt.Expression.(*ast.IfExpression)

		if tt.exprect != stmte.String() {
			t.Errorf("Exprected=%s, Got=%s", tt.exprect, stmte.String())
		}
	}
}

func TestCallBlock(t *testing.T) {
	test := []struct {
		input    string
		expected string
	}{
		{`?careFile`, "careFile"},
	}

	for _, tt := range test {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()

		stmt := program.Statements[0].(*ast.ExpressionStatement)

		stmte := stmt.Expression.(*ast.CallBlockExpression)

		if tt.expected != stmte.String() {
			t.Errorf("Expected=%s, got=%s", tt.expected, stmte.String())
		}
	}
}

func TestInfixExpression(t *testing.T) {
	test := []struct {
		input    string
		expected string
	}{
		{"3 + 4", "(3 + 4)"},
		{"3 > 5 == false", "((3 > 5) == false)"},
		{"(5 + 5) * 2", "((5 + 5) * 2)"},
	}

	for _, tt := range test {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()

		stmt := program.Statements[0].(*ast.ExpressionStatement)

		stmte := stmt.Expression.(*ast.InfixExpression)

		if tt.expected != stmte.String() {
			t.Errorf("expectecd=%s, got=%s", tt.expected, stmte.String())
		}
	}
}

func TestBlock(t *testing.T) {
	test := []struct {
		input  string
		expect string
	}{
		{"Block createFile{return 23fsd;}", "order"},
	}

	for _, tt := range test {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		stmt := program.Statements[0].(*ast.ExpressionStatement)

		if stmt.Token.Type != token.BLOCK {
			t.Errorf("No es tipo blocke")
		}

		stmte := stmt.Expression.(*ast.BlockLiteral)
		stdec := stmte.Body.Statements[0].(*ast.DeclareStatement)
		if stdec.Name.Value != tt.expect {
			t.Errorf("El nombre del bloque no ah sido el esperado")
		}
	}
}

func TestReturnStatement(t *testing.T) {
	test := []struct {
		input  string
		expect string
	}{
		{"return 23;", "return"},
	}

	for _, tt := range test {
		l := lexer.New(tt.input)
		p := New(l)

		program := p.ParseProgram()

		stmt := program.Statements[0].(*ast.ReturnStatement)

		if tt.expect != stmt.TokenLiteral() {
			t.Errorf("Expect=%s, got=%s", tt.expect, stmt.TokenLiteral())
		}
	}
}

func TestDecStatement(t *testing.T) {
	test := []struct {
		input  string
		expect string
	}{
		{"dec xd int -> 324;", "xd"},
		{`dec stringi string -> "fssfg";`, "stringi"},
	}
	for _, tt := range test {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		sttm := program.Statements[0].(*ast.DeclareStatement)

		if sttm.Name.Value != tt.expect {
			t.Errorf("Not edpected")
		}
	}
}

// PARSE ERRORS CHEKER
func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("Parser error: %q", msg)
	}
	t.FailNow()
}
