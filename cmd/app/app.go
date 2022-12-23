package main

import (
	"fmt"
	"go-vmachine/pkgs/vm"
	"go-vmachine/pkgs/vm/instruction"
)

func main() {
	println("Simple Virtual Machine on GoLang v1.0")
	vmachine := vm.NewVM()

	instructions := []*instruction.Instruction{
		instruction.NewNumber(20),
		instruction.NewNumber(30),
		instruction.NewOperator(instruction.OpCodeAdd),
		instruction.NewNumber(10),
		instruction.NewOperator(instruction.OpCodeSub),
	}

	machine_result := vmachine.Run(instructions...)
	fmt.Printf("Result: %v", machine_result)
}
