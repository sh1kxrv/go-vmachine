package main

import (
	"go-vmachine/pkgs/vm/compiler"
	"go-vmachine/pkgs/vm/cpu"
	"go-vmachine/pkgs/vm/lexer"
)

func main() {
	code := `
		# add program
		add 0x10 0x10
		add 0x1 0x10
		sub 0x10 0x12
	`
	lexer := lexer.NewLexer(code)
	compiler := compiler.NewCompiler(lexer)
	cpu := cpu.NewCPU()
	compiled := compiler.Compile()
	cpu.Load(compiled)
	cpu.Run()
}
