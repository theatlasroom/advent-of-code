package intcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerformOperation(t *testing.T) {
	instruction := Intcode{
		opcode: OcADD,
		a:      1,
		b:      2,
		output: 3,
	}
	result := PerformOperation(instruction, []int{0, 10, 20, 3})
	assert.Equal(t, result[3], 30)

	instruction.opcode = OcMULTI
	result = PerformOperation(instruction, []int{0, 10, 20, 3})
	assert.Equal(t, result[3], 200)
	instruction.opcode = OcMULTI

	instruction.opcode = 10
	result = PerformOperation(instruction, []int{0, 10, 20, 3})
	assert.Equal(t, result[3], 3)
}

func TestComputeNextProgram(t *testing.T) {
	program := []int{1, 0, 0, 0, 99}
	result := []int{2, 0, 0, 0, 99}
	computed := ComputeNextProgram(program)
	assert.Equal(t, computed, result)

	program = []int{2, 3, 0, 3, 99}
	result = []int{2, 3, 0, 6, 99}
	computed = ComputeNextProgram(program)
	assert.Equal(t, computed, result)

	program = []int{2, 4, 4, 5, 99, 0}
	result = []int{2, 4, 4, 5, 99, 9801}
	computed = ComputeNextProgram(program)
	assert.Equal(t, computed, result)

	program = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	result = []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	computed = ComputeNextProgram(program)
	assert.Equal(t, computed, result)
}
