package gs_test

import (
	"testing"

	"github.com/ek220/gs"
	"github.com/stretchr/testify/assert"
)

func TestNewStackIsEmpty(t *testing.T) {
	assert.True(t, gs.NewStack[struct{}]().Empty())
}

func TestPutTwoToStackAndCheckLen(t *testing.T) {
	s := gs.NewStack[string]()
	s.Push("aaa")
	s.Push("bbb")
	assert.Equal(t, 2, s.Len())
}

func TestPutTwoToStackAndInReverseOrder(t *testing.T) {
	s := gs.NewStack[string]()
	s.Push("aaa")
	s.Push("bbb")
	assert.Equal(t, "bbb", s.Pop())
	assert.Equal(t, "aaa", s.Pop())
	assert.True(t, s.Empty())
}
