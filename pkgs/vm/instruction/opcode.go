package instruction

type OpCode uint8

const (
	OpCodeNop OpCode = iota

	// Operator
	OpCodeOperator
	OpCodeAdd
	OpCodeSub
	OpCodeMul
	OpCodeDiv

	// Const
	OpCodeLdci1 // int8
	OpCodeLdci2 // int16
	OpCodeLdci4 // int32
	OpCodeLdci8 // int64
	OpCodeLdcf4 // float32
	OpCodeLdcf8 // float64

	// Mecha
	OpCodeRet
)

func (o *OpCode) String() string {
	switch *o {
	case OpCodeNop:
		return "nop"
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
	case OpCodeLdci4:
		return "ldci4"
	case OpCodeLdci8:
		return "ldci8"
	case OpCodeLdcf4:
		return "ldcf4"
	case OpCodeLdcf8:
		return "ldcf8"
	case OpCodeRet:
		return "ret"
	default:
		return "unknown"
	}
}
