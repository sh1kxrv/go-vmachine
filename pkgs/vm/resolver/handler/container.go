package handler

import (
	"go-vmachine/pkgs/vm/instruction"
	"go-vmachine/pkgs/vm/stack"
)

/*

Container:
ldci.1 100
ldci.1 10
container 0x41 ldci.2 # transformer
operator mul 0x41 # -> 100 * 10 = 1000 -> into 0x41 container

*/

func ContainerHandler(s *stack.Stack, instr *instruction.Instruction) {

}
