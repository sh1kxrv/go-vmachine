package handler

import (
	"fmt"
	"go-vmachine/pkgs/vm/instruction"
)

func AnalyseOperator(current *instruction.Instruction, list instruction.InstructionList) {
	v1, v2 := list.Peek(current.Offset - 1), list.Peek(current.Offset - 2)
	println(fmt.Sprintf("AnalyseOperator: %v %v",v1, v2))
}
