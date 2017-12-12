package utils

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberLabelsToStr(t *testing.T) {
	i1 := []int{1, 2, 3, 4, 5}
	s1 := []string{"1", "2", "3", "4", "5"}
	assert.True(t, reflect.DeepEqual(s1, NumberLabelsToStr(i1)))
}
