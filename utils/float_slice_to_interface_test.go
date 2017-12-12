package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloatSliceToInterface(t *testing.T) {
	f641 := []float64{1, 2, 3, 4, 5}
	i1 := []interface{}{1.0, 2.0, 3.0, 4.0, 5.0}
	assert.Equal(t, i1, FloatSliceToInterface(f641))
}
