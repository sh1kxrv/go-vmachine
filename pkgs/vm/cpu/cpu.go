package cpu

type CPU struct {
	Stack   *Stack
	Program [0xFFFFFF]byte
}

func NewCPU() *CPU {
	return &CPU{
		Stack: NewStack(),
	}
}

func (c *CPU) Resolve() {
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
