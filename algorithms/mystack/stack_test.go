package mystack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Stack(t *testing.T) {
	assert := assert.New(t)

	s := NewStack()
	assert.True(s.IsEmpty(), "Check s is empty or not")

	start, end := 1, 100
	for i := start; i <= end; i++ {
		s.Push(i)
		assert.Equal(i-start+1, s.Len(), "Check s.Len() after push")
	}

	for i := end; i >= start; i-- {
		assert.Equal(i, s.Pop(), "Pop from s")
	}
	assert.True(s.IsEmpty(), "Check s is empty or not after pop")
}