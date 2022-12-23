package handler

import (
	"fmt"
	"go-vmachine/pkgs/vm/instruction"
)

func AnalyseOperator(instr *instruction.Instruction) {
	println(fmt.Sprintf("AnalyseOperator: %v", instr))
}
