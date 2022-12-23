package analyser

import (
	"go-vmachine/pkgs/vm/analyser/handler"
	"go-vmachine/pkgs/vm/instruction"
)

type AnalyserHandler func(*instruction.Instruction)

type Analyser struct {
	Analysers map[instruction.OpCode]AnalyserHandler
}

func NewAnalyser() *Analyser {
	return &Analyser{
		Analysers: map[instruction.OpCode]AnalyserHandler {
			instruction.OpCodeLdci1: handler.AnalyseLdci1,
			instruction.OpCodeLdci2: handler.AnalyseLdci2,
			instruction.OpCodeLdci4: handler.AnalyseLdci4,
			instruction.OpCodeLdci8: handler.AnalyseLdci8,
			instruction.OpCodeLdcf4: handler.AnalyseLdcf4,
			instruction.OpCodeLdcf8: handler.AnalyseLdcf8,
		},
	}
}

func (a *Analyser) Analyse(instr *instruction.Instruction){
	a.Analysers[instr.OpCode](instr)
}
