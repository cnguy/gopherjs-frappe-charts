package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBtoi(t *testing.T) {
	assert.Equal(t, Btoi(true), 1)
	assert.Equal(t, Btoi(false), 0)
}
