package lexer

import (
	"fmt"
	"testing"

	"github.com/Hydrogen/src/token"
)

func TestNextToken(t *testing.T) {
	input := `
	dec steep int -> 34;
	dec net string -> "Luis";
	dec bol bool -> true;
	fn main(x int) <- int {
		return 23
	};
	##Hola mundo

	#createFile{
		return 23;
	}

	#def digit 23

	if (3 < 5) {
		?createFile
	}else{
		?createFile
	}

	Block createFile {
		fn xd(x string) <- string {
			return "sdffsdfsd";
		}
	}

	?createFile
	`

	test := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.DEC, "dec"},
		{token.IDENT, "steep"},
		{token.INT, "int"},
		{token.ASSIGN, "->"},
		{token.INT, "34"},
		{token.SEMICOLON, ";"},
		{token.DEC, "dec"},
		{token.IDENT, "net"},
		{token.STRING, "string"},
		{token.ASSIGN, "->"},
		{token.STRING, "Luis"},
		{token.SEMICOLON, ";"},
		{token.DEC, "dec"},
		{token.IDENT, "bol"},
		{token.BOOL, "bool"},
		{token.ASSIGN, "->"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.FUNCTION, "fn"},
		{token.IDENT, "main"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.INT, "int"},
		{token.RPAREN, ")"},
		{token.DEVOLUCION, "<-"},
		{token.INT, "int"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.INT, "23"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.COMMENT, "Hola mundo"},
		{token.HASH, "#"},
		{token.IDENT, "createFile"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.INT, "23"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.HASH, "#"},
		{token.DEF, "def"},
		{token.IDENT, "digit"},
		{token.INT, "23"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "3"},
		{token.LT, "<"},
		{token.INT, "5"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.INTERROG, "?"},
		{token.IDENT, "createFile"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.INTERROG, "?"},
		{token.IDENT, "createFile"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range test {
		tok := l.NextToken()
		fmt.Println(tok)

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] -tokentype wrong expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - Literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
