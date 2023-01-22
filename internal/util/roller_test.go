package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Roll_d20(t *testing.T) {
	for i := 0; i < 25; i++ {
		roll := Roll(20)
		assert.Greater(t, roll, 0)
		assert.LessOrEqual(t, roll, 20)
	}
}

func Test_Roll_d6(t *testing.T) {
	for i := 0; i < 25; i++ {
		roll := Roll(6)
		assert.Greater(t, roll, 0)
		assert.LessOrEqual(t, roll, 6)
	}
}

func Test_Roll_d12(t *testing.T) {
	for i := 0; i < 25; i++ {
		roll := Roll(12)
		assert.Greater(t, roll, 0)
		assert.LessOrEqual(t, roll, 12)
	}
}
