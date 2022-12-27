package cpu

import "go-vmachine/pkgs/vm/opcode"

type CPU struct {
	Stack   *Stack
	Program [0xFFFFFF]byte
	Pointer int
}

func NewCPU() *CPU {
	return &CPU{
		Stack:   NewStack(),
		Pointer: 0,
	}
}

func (c *CPU) Run() {
	flag := true
	for flag {
		op := opcode.NewOpCode(c.Program[c.Pointer])
		switch int(op.V()) {
		case opcode.EXIT:
			flag = false
		case opcode.ADD:
			c.Pointer++
			// todo: read add operation
			// read bytes to int
		}
	}
}

func (c *CPU) Load(program []byte) {
	if len(program) > 0xFFFFFF {
		panic("Program is too big")
	}
	for i, b := range program {
		// lint copy() silence
		c.Program[i+0] = b
	}
}
