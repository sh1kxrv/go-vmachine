package compiler

import (
	"go-vmachine/pkgs/vm/opcode"
	"go-vmachine/pkgs/vm/token"
)

func TokenToOpcode(t token.Token) byte {
	switch t.Type {
	case token.ADD:
		return opcode.ADD
	case token.SUB:
		return opcode.SUB
	case token.MUL:
		return opcode.MUL
	case token.DIV:
		return opcode.DIV
	default:
		panic("Unknown token type")
	}
}
