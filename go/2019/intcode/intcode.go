package intcode

// Opcodes indicating the various valid operations
const (
	OcHALT  int = 99
	OcADD   int = 1
	OcMULTI int = 2
)

// Intcode struct type
type Intcode struct {
	opcode       int
	a, b, output int
}

// ComputeNextProgram halts the running program
func ComputeNextProgram(program []int) []int {
	offset := 4
	position := 0
	halted := false
	eof := false

	np := make([]int, len(program))
	copy(np, program)

	for !halted && !eof {
		if np[position] != OcHALT {
			instruction := Intcode{
				opcode: np[position],
				a:      np[position+1],
				b:      np[position+2],
				output: np[position+3],
			}

			np = PerformOperation(instruction, np)
			position += offset
		} else {
			halted = true
			break
		}

		if position >= len(program) {
			eof = true
		}
		if position >= len(program) {
			eof = true
		}
	}
	return np
}

// PerformOperation checks the opcode and applies the operation to the input program
func PerformOperation(instruction Intcode, program []int) []int {
	switch instruction.opcode {
	case OcADD:
		program[instruction.output] = program[instruction.a] + program[instruction.b]
	case OcMULTI:
		program[instruction.output] = program[instruction.a] * program[instruction.b]
	default:
		// noop
	}
	return program
}
