package main

import (
	"go-vmachine/internal/lexer"
	"go-vmachine/internal/token"
)

func main() {
	code := `
		# test comment
		0x10
		1337
		"hello!
	`
	lexer := lexer.NewLexer(code)
	for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
		println(tok.Type, tok.Literal)
	}
}
