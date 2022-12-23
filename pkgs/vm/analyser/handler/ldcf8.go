package handler

import (
	"fmt"
	"go-vmachine/pkgs/vm/instruction"
)

func AnalyseLdcf8(instr *instruction.Instruction)  {
	println(fmt.Sprintf("AnalyseLdcf8: %v", instr))
}
