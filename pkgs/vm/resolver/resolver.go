package resolver

import (
	"fmt"
	"go-vmachine/pkgs/vm/instruction"
	"go-vmachine/pkgs/vm/stack"
)

type ResolverHandler func(*stack.Stack, *instruction.Instruction)

type Resolver struct {
	Handlers map[instruction.OpCode]ResolverHandler
}

func NewResolver() *Resolver {
	return &Resolver{
		Handlers: make(map[instruction.OpCode]ResolverHandler),
	}
}

func (r *Resolver) Register(opCode instruction.OpCode, handler ResolverHandler) *Resolver {
	r.Handlers[opCode] = handler
	return r
}

func (r *Resolver) Resolve(stack *stack.Stack, instruction *instruction.Instruction) {
	handler_to_resolve := r.Handlers[instruction.OpCode]
	if handler_to_resolve == nil {
		panic(fmt.Sprintf("No handler found for instruction %v", instruction.OpCode))
	}
	handler_to_resolve(stack, instruction)
}
