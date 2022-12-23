package vm

import (
	"go-vmachine/pkgs/logger"
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

	if len(vm.Stack.StackData) > 1 {
		logger.Warn("Stack Data Length > 1 instruction after execution")
	}

	ret := vm.Stack.Pop()
	if ret != nil {
		return ret
	}
	return nil
}
