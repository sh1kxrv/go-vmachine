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

func (vm *VM) Run(instructions ...*instruction.Instruction) interface{} {
	analyser := analyser.NewAnalyser()
	resolver := resolver.NewResolver()
	transformer := transformer.NewTransformer()

	analyser.Analyse(instructions)
	transformer.Transform(instructions)
	resolver.Resolve(instructions, vm.Stack)

	return vm.Stack.Pop()
}
