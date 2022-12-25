package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, liternal string) Token {
	return Token{Type: tokenType, Literal: liternal}
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	IDENT   = "IDENT"
	INT     = "INT"
	STRING  = "STRING"
	COMMA   = "COMMA"

	// Math
	ADD = "ADD"
	SUB = "SUB"
	MUL = "MUL"
	DIV = "DIV"
)

var keywords = map[string]TokenType{
	"add": ADD,
	"sub": SUB,
	"mul": MUL,
	"div": DIV,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
