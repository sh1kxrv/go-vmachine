package main

import (
	"go-vmachine/pkgs/vm/compiler"
	"go-vmachine/pkgs/vm/lexer"
)

func main() {
	code := `
		# test comment
		0x10
		10.1
	`
	lexer := lexer.NewLexer(code)
	compiler := compiler.NewCompiler(lexer)
	compiler.Compile()
}
