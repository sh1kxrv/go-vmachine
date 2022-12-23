package transformer

import (
	"go-vmachine/pkgs/logger"
	"go-vmachine/pkgs/vm/instruction"
	"go-vmachine/pkgs/vm/transformer/handler"
)

type TransformerHandler func(*instruction.Instruction)

type Transformer struct {
	Transformers map[instruction.OpCode]TransformerHandler
}

func NewTransformer() *Transformer {
	return &Transformer{
		Transformers: map[instruction.OpCode]TransformerHandler{
			instruction.OpCodeLdci1: handler.TransformLdci1,
			instruction.OpCodeLdci2: handler.TransformLdci2,
			instruction.OpCodeLdci4: handler.TransformLdci4,
			instruction.OpCodeLdci8: handler.TransformLdci8,
			instruction.OpCodeLdcf4: handler.TransformLdcf4,
			instruction.OpCodeLdcf8: handler.TransformLdcf8,
		},
	}
}

// Instructions optimizer
func (t *Transformer) Transform(instructions []*instruction.Instruction) {
	for _, i := range instructions {
		handler := t.Transformers[i.OpCode]
		if handler == nil {
			logger.Warn("No transform handler found for instruction", i)
			continue
		}
		handler(i)
	}
}
