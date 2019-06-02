package opcode

var (
	EXIT      = 0x00
	INT_STORE = 0x01
	INT_PRINT = 0x02
	ADD_OP    = 0x03
	SUB_OP    = 0x04
	MUL_OP    = 0x05
	DIV_OP    = 0x06
	INC_OP    = 0x07
	DES_OP    = 0x08
)

type OpCode struct {
	instruction byte
}

func NewOpCode(instruction byte) *OpCode {
	opcode := &OpCode{}
	opcode.instruction = instruction
	return opcode
}

func (opcode *OpCode) toString() string {
	switch int(opcode.instruction) {
	case EXIT:
		return "exit"

	case INT_STORE:
		return "INT_STORE"

	case INT_PRINT:
		return "INT_PRINT"

	case ADD_OP:
		return "ADD_OP"

	case SUB_OP:
		return "SUB_OP"

	case MUL_OP:
		return "MUL_OP"

	case DIV_OP:
		return "DIV_OP"

	case INC_OP:
		return "INC_OP"

	case DES_OP:
		return "DES_OP"
	}
	return "Undefined OPCODE"
}

func (opcode *OpCode) toByteCode() byte {
	return (opcode.instruction)
}
