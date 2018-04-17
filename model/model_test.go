package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateRange(t *testing.T) {
	assert.NoError(t, validate.Var([]float64{0, 1}, "range"))
	assert.NoError(t, validate.Var([2]float64{0, 1}, "range"))
	assert.Error(t, validate.Var([]float64{1, 1}, "range"))
	assert.Error(t, validate.Var([]float64{2, 1}, "range"))
}
