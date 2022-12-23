package vm

import (
	"go-vmachine/pkgs/vm/analyser"
	"go-vmachine/pkgs/vm/instruction"
	"go-vmachine/pkgs/vm/resolver"
	"go-vmachine/pkgs/vm/resolver/handler"
	"go-vmachine/pkgs/vm/stack"
)

type VM struct {
	Stack *stack.Stack
}

func NewVM() *VM {
	return &VM{
		Stack: stack.NewStack(),
	}
}

func (vm *VM) Run() interface{} {
	analyser := analyser.NewAnalyser()
	resolver := resolver.NewResolver()
	resolver.
		Register(instruction.OpCodeLdci1, handler.NumberHandler).
		Register(instruction.OpCodeLdci2, handler.NumberHandler).
		Register(instruction.OpCodeLdci4, handler.NumberHandler).
		Register(instruction.OpCodeLdci8, handler.NumberHandler).
		Register(instruction.OpCodeLdcf4, handler.NumberHandler).
		Register(instruction.OpCodeLdcf8, handler.NumberHandler).
		Register(instruction.OpCodeOperator, handler.OperatorHandler).
		Register(instruction.OpCodeRet, handler.ReturnHandler)

	for _, instruction := range vm.Stack.Instructions {
		analyser.Analyse(instruction)
		resolver.Resolve(vm.Stack, instruction)
	}
	return vm.Stack.Pop()
}

func (vm *VM) Push(instructions ...*instruction.Instruction) *VM {
	vm.Stack.Push(instructions...)
	return vm
}
