package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberLabelsToStr(t *testing.T) {
	i1 := []int{1, 2, 3, 4, 5}
	s1 := []string{"1", "2", "3", "4", "5"}
	assert.Equal(t, NumberLabelsToStr(i1), s1)
}
