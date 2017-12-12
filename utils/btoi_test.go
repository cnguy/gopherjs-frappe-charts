package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBtoi(t *testing.T) {
	assert.Equal(t, 1, Btoi(true))
	assert.Equal(t, 0, Btoi(false))
}
