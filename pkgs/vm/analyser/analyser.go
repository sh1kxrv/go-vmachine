package analyser

import (
	"go-vmachine/pkgs/logger"
	"go-vmachine/pkgs/vm/analyser/handler"
	"go-vmachine/pkgs/vm/instruction"
)

type AnalyserHandler func(*instruction.Instruction, instruction.InstructionList)

type Analyser struct {
	Analysers map[instruction.OpCode]AnalyserHandler
}

func NewAnalyser() *Analyser {
	return &Analyser{
		Analysers: map[instruction.OpCode]AnalyserHandler {
			instruction.OpCodeOperator: handler.AnalyseOperator,
		},
	}
}

// ...
func (a *Analyser) Analyse(instructions instruction.InstructionList){
	// Calculate offsets
	for i, v := range instructions {
		v.SetOffset(i)
	}

	for _, i := range instructions {
		handler := a.Analysers[i.OpCode]
		if handler == nil {
			logger.Warn("No analyse handler found for instruction", i)
			continue
		}
		handler(i, instructions)
	}
}
