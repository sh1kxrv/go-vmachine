package main

import (
	"fmt"
	"go-vmachine/pkgs/vm"
	"go-vmachine/pkgs/vm/instruction"
)

func main() {
	println("Simple Virtual Machine on GoLang v1.0")
	vmachine := vm.NewVM()
	vmachine.Push(
		instruction.NewNumber(20),
		instruction.NewNumber(30),
		instruction.NewOperator(instruction.OpCodeAdd),
		instruction.NewNumber(10),
		instruction.NewOperator(instruction.OpCodeSub),
	)
	vmachine.Stack.Debug()
	machine_result := vmachine.Run()
	fmt.Printf("Result: %v", machine_result)
}
