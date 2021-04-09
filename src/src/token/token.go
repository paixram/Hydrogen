package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (

	// IDENTIFIERS + LITERALS
	INT    = "INT"
	STRING = "STRING"
	BOOL   = "BOOL"
	FLOAT  = "FLOAT"
	IDENT  = "IDENT"

	// Operators
	ASSIGN     = "->"
	DEVOLUCION = "<-"
	PLUS       = "+"
	MINUS      = "-"

	LT       = "<"
	GT       = ">"
	SLASH    = "/"
	ASTERISK = "*"

	// DELIMITERS
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	RBRACE = "}"
	LBRACE = "{"

	// KEYWORDS
	FUNCTION = "FUNCTION"
	DEC      = "DEC"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	DEF      = "DEF"
	BLOCK    = "BLOCK"
	STOP     = "STOP"

	EQ     = "=="
	NOT_EQ = "!="

	HASH     = "#"
	INTERROG = "?"

	// ILLEGAL + EOF
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// COMMENT
	COMMENT = "COMMENT"

	BANG  = "!"
	TRUE  = "TRUE"
	FALSE = "FALSE"

	// HANDLER TOKENS
	TYPE = "TYPE"

	// TYPES
	TYPEINT    = "INTTYPE"
	TYPESTRING = "STRINGTYPE"
	TYPEBOOL   = "BOOLTYPE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"dec":    DEC,
	"return": RETURN,
	"int":    INT,
	"def":    DEF,
	"else":   ELSE,
	"if":     IF,
	"string": STRING,
	"bool":   BOOL,
	"true":   TRUE,
	"false":  FALSE,
	"Block":  BLOCK,
	"?":      INTERROG,
	"STOP":   STOP,
	//"?":      INTERROG,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
