package instruction

type OpCode uint8

const (
	OpCodeNop      OpCode = iota
	OpCodeNumber   OpCode = 0x1
	OpCodeString   OpCode = 0x2
	OpCodeOperator OpCode = 0x3
	OpCodeAdd      OpCode = 0x4
	OpCodeSub      OpCode = 0x5
	OpCodeMul      OpCode = 0x6
	OpCodeDiv      OpCode = 0x7
	OpCodeRet      OpCode = 0x8
)

func (o *OpCode) String() string {
	switch *o {
	case OpCodeNop:
		return "nop"
	case OpCodeNumber:
		return "number"
	case OpCodeString:
		return "string"
	case OpCodeOperator:
		return "operator"
	case OpCodeAdd:
		return "add"
	case OpCodeSub:
		return "sub"
	case OpCodeMul:
		return "mul"
	case OpCodeDiv:
		return "div"
	case OpCodeRet:
		return "ret"
	default:
		return "unknown"
	}
}
