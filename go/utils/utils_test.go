package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrToIntArr(t *testing.T) {
	input := "1,1,1,4,99,5,6,0,99"
	res := StrToIntArr(input)

	assert.Equal(t, res, []int{1, 1, 1, 4, 99, 5, 6, 0, 99})
}
