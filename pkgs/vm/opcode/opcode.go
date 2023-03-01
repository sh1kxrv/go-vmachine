package opcode

type OpCode struct {
	Instruction byte
}

func NewOpCode(instruction byte) *OpCode {
	return &OpCode{
		Instruction: instruction,
	}
}

func (o *OpCode) V() byte {
	return o.Instruction
}

const (
	EXIT = 0x00

	// Math opcodes
	ADD = 0x01
	SUB = 0x02
	MUL = 0x03
	DIV = 0x04

	// Stack opcodes
	PUSH = 0x10
	POP  = 0x11
	PEEK = 0x12

	// Read
	READ = 0x13

	// Variables
	REG = 0xF

	NOP = 0xFF
)

var opcodes = map[string]byte{
	"UNKNOWN": 0x00,
	"ADD":     ADD,
	"SUB":     SUB,
	"MUL":     MUL,
	"DIV":     DIV,
	"REG":     REG,
	"READ":    READ,
}

func LookupOpCode(token string) byte {
	if opcode, ok := opcodes[token]; ok {
		return opcode
	}
	return NOP
}
