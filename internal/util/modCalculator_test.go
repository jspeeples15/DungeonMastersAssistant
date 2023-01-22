package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModCalculator_CalculatorModifier(t *testing.T) {
	assert.Equal(t, 1, CalculateModifier(12))
	assert.Equal(t, 1, CalculateModifier(13))
	assert.Equal(t, 0, CalculateModifier(11))
	assert.Equal(t, 0, CalculateModifier(9))
	assert.Equal(t, -1, CalculateModifier(8))
	assert.Equal(t, 3, CalculateModifier(16))
	assert.Equal(t, -1, CalculateModifier(7))
	assert.Equal(t, -2, CalculateModifier(6))
}
