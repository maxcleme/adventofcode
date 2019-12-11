package utils

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrUnknownOpCode = Error("unknown opcode")

	OpcodeAdd  = 1
	OpcodeMult = 2
	OpcodeExit = 99
)

func Run(instructions []int) ([]int, error) {
	// clone instructions slice
	copy := make([]int, len(instructions))
	for i, instruction := range instructions {
		copy[i] = instruction
	}

	i := 0
	for i < len(copy) {
		opcode := copy[i]
		if opcode == OpcodeExit {
			break
		}
		val1 := copy[i+1]
		val2 := copy[i+2]
		dest := copy[i+3]

		res := 0
		switch opcode {
		case OpcodeAdd:
			res = copy[val1] + copy[val2]
		case OpcodeMult:
			res = copy[val1] * copy[val2]
		default:
			return nil, ErrUnknownOpCode
		}

		copy[dest] = res
		i += 4
	}
	return copy, nil
}
