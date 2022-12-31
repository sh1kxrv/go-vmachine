package cpu

import (
	"fmt"
	"go-vmachine/pkgs/vm/opcode"
)

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

func (c *CPU) ReadValue() int {
	l := int(c.NextByteInc())
	l2 := int(c.NextByteInc())
	return l + l2*256
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
			v := c.ReadValue()
			v2 := c.ReadValue()
			c.Stack.Push(v + v2)
			println(fmt.Printf("%v + %v", v, v2))
		case opcode.SUB:
			c.Pointer++
			v := c.ReadValue()
			v2 := c.ReadValue()
			c.Stack.Push(v - v2)
			println(fmt.Printf("%v - %v", v, v2))
		default:
			panic("wtf")
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

func (c *CPU) CurrentByte() byte {
	return c.Program[c.Pointer]
}

func (c *CPU) NextByteInc() byte {
	b := c.Program[c.Pointer]
	c.Pointer++
	return b
}
