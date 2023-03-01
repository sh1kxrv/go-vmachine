package cpu

import (
	"encoding/binary"
	"fmt"
	"go-vmachine/pkgs/vm/opcode"
)

type CPU struct {
	Stack   *Stack
	Program [0xFFFFFF]byte
	Regs    [0xFF]byte
	Pointer int
}

func NewCPU() *CPU {
	return &CPU{
		Stack:   NewStack(),
		Pointer: 0,
		Regs:    [0xFF]byte{},
	}
}

func (c *CPU) ReadValue() int64 {
	intLen := int(c.NextByte())
	readed := make([]byte, 8)
	for i := 0; i < intLen; i++ {
		readed[i] = c.NextByte()
	}
	return int64(binary.LittleEndian.Uint64(readed))
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
			v3 := c.NextByte()
			if v3 == opcode.NOP {
				c.Stack.Push(v + v2)
			} else if v3 < 0xFF {
				// ...
			} else {
				panic("Register should be less 0xFF")
			}
			println(fmt.Sprintf("%d + %d -> %d", v, v2, v+v2))
		case opcode.SUB:
			c.Pointer++
			v := c.ReadValue()
			v2 := c.ReadValue()
			c.Stack.Push(v - v2)
			println(fmt.Sprintf("%d - %d -> %d", v, v2, v-v2))
		case opcode.READ:
			c.Pointer++
			// ...
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

func (c *CPU) NextByte() byte {
	b := c.Program[c.Pointer]
	c.Pointer++
	return b
}
