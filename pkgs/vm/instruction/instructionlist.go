package instruction

type InstructionList []*Instruction

func (l InstructionList) Len() int {
	return len(l)
}

func (l InstructionList) Peek(offset int) *Instruction {
	if offset < 0 || offset >= len(l) {
		return nil
	}
	return l[offset]
}
