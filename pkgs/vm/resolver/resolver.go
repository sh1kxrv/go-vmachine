package resolver

import (
	"fmt"
	"go-vmachine/pkgs/vm/instruction"
	"go-vmachine/pkgs/vm/resolver/handler"
	"go-vmachine/pkgs/vm/stack"
)

type ResolverHandler func(*stack.Stack, *instruction.Instruction)

type Resolver struct {
	Handlers map[instruction.OpCode]ResolverHandler
}

func NewResolver() *Resolver {
	return &Resolver{
		Handlers: map[instruction.OpCode]ResolverHandler {
			instruction.OpCodeLdci1: handler.NumberHandler,
			instruction.OpCodeLdci2: handler.NumberHandler,
			instruction.OpCodeLdci4: handler.NumberHandler,
			instruction.OpCodeLdci8: handler.NumberHandler,
			instruction.OpCodeLdcf4: handler.NumberHandler,
			instruction.OpCodeLdcf8: handler.NumberHandler,
			instruction.OpCodeOperator: handler.OperatorHandler,
			instruction.OpCodeRet: handler.ReturnHandler,
		},
	}
}


func (r *Resolver) Resolve(instructions []*instruction.Instruction, stack *stack.Stack) {
	for _, instruction := range instructions {
		handler := r.Handlers[instruction.OpCode]
		if handler == nil {
			panic(fmt.Sprintf("No handler found for instruction %v", instruction))
		}
		handler(stack, instruction)
	}
}
