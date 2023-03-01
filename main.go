package main

import (
	"go-vmachine/pkgs/vm/compiler"
	"go-vmachine/pkgs/vm/cpu"
	"go-vmachine/pkgs/vm/lexer"
)

func main() {
	code := `
		reg 0x01 0x00 					# set register 0x01 with default value inner - 0x00
		add 0x10 0x10 to 0x01   # 16 + 16 and write it to register with index 0x01
		read 0x01               # read from 0x01
	`
	lexer := lexer.NewLexer(code)
	compiler := compiler.NewCompiler(lexer)
	cpu := cpu.NewCPU()

	compiled := compiler.Compile()
	cpu.Load(compiled)
	cpu.Run()
}
