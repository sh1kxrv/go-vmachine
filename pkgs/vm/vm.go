package vm

import (
	"go-vmachine/pkgs/vm/analyser"
	"go-vmachine/pkgs/vm/instruction"
	"go-vmachine/pkgs/vm/resolver"
	"go-vmachine/pkgs/vm/stack"
	"go-vmachine/pkgs/vm/transformer"
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
	transformer := transformer.NewTransformer()

	analyser.Analyse(vm.Stack.Instructions)
	transformer.Transform(vm.Stack.Instructions)
	resolver.Resolve(vm.Stack)

	return vm.Stack.Pop()
}

func (vm *VM) Push(instructions ...*instruction.Instruction) *VM {
	vm.Stack.Push(instructions...)
	return vm
}
