package gs_test

import (
	"sort"
	"testing"

	"github.com/ek220/gs"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	s := gs.NewSet[string]().Add("foo").Add("bar").Add("asd").Remove("bar")

	assert.Equal(t, 2, s.Len())
	assert.True(t, s.Contains("foo"))
	assert.True(t, s.Contains("asd"))
	assert.False(t, s.Contains("bar"))
	assert.False(t, s.Contains("baz"))

	s.Add("baz")
	assert.True(t, s.Contains("baz"))
}

func TestSetValues(t *testing.T) {
	s := gs.NewSet[string]().Add("foo").Add("bar").Add("asd").Remove("bar")

	values := s.Values()
	s.Remove("foo")

	sort.Strings(values)
	assert.Equal(t, []string{"asd", "foo"}, values)
}
