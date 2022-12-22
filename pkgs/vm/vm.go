package vm

import (
	"go-vmachine/pkgs/vm/handler"
	"go-vmachine/pkgs/vm/instruction"
	"go-vmachine/pkgs/vm/resolver"
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
	resolver := resolver.NewResolver()
	resolver.
		Register(instruction.OpCodeNumber, handler.NumberHandler).
		Register(instruction.OpCodeString, handler.OperatorHandler).
		Register(instruction.OpCodeOperator, handler.OperatorHandler).
		Register(instruction.OpCodeRet, handler.ReturnHandler)

	for _, instruction := range vm.Stack.Instructions {
		resolver.Resolve(vm.Stack, instruction)
	}
	return vm.Stack.Pop()
}

func (vm *VM) Push(instructions ...*instruction.Instruction) *VM {
	vm.Stack.Push(instructions...)
	return vm
}
