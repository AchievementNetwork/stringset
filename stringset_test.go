package stringset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	ss := New()
	assert.Equal(t, 0, ss.Length())
	ss.Add("this")
	assert.Equal(t, 1, ss.Length())
	ss.Add("is", "a", "test")
	assert.Equal(t, 4, ss.Length())
	assert.True(t, ss.Contains("this"))
	assert.True(t, ss.Contains("is"))
	assert.True(t, ss.Contains("a"))
	assert.True(t, ss.Contains("test"))
	assert.False(t, ss.Contains("that"))
	ss.Add("a")
	assert.Equal(t, 4, ss.Length())
	checkEquivalence(t, ss.Strings(), []string{"this", "is", "a", "test"})
	ss.Delete("this")
	checkEquivalence(t, ss.Strings(), []string{"is", "a", "test"})
	ss.Delete("is", "nothing")
	checkEquivalence(t, ss.Strings(), []string{"a", "test"})
	ss.Add("this", "is", "is", "a", "test")
	checkEquivalence(t, ss.Strings(), []string{"this", "is", "a", "test"})
}

func TestEqual(t *testing.T) {
	ss1 := New()
	ss2 := New()
	assert.True(t, ss1.Equals(ss2))
	assert.True(t, ss2.Equals(ss1))
	ss1.Add("a")
	assert.False(t, ss1.Equals(ss2))
	assert.False(t, ss2.Equals(ss1))
	ss2.Add("a")
	assert.True(t, ss1.Equals(ss2))
	assert.True(t, ss2.Equals(ss1))
	ss2.Add("a")
	assert.True(t, ss1.Equals(ss2))
	assert.True(t, ss2.Equals(ss1))
	ss1.Add("b", "c")
	ss2.Add("b", "c")
	assert.True(t, ss1.Equals(ss2))
	ss1.Add("d")
	ss2.Add("e")
	assert.False(t, ss1.Equals(ss2))
}

func TestOperations(t *testing.T) {
	ss1 := New().Add("this", "is", "a", "test")
	ss2 := New().Add("this", "was", "an", "interesting", "test")

	inter := ss1.Intersection(ss2)
	assert.True(t, inter.Equals(New().Add("this", "test")))
	union := ss1.Union(ss2)
	assert.True(t, union.Equals(New().Add("this", "was", "an", "interesting", "test", "is", "a")))
	s1 := ss1.Difference(ss2)
	assert.True(t, s1.Equals(New().Add("is", "a")))
	s2 := ss2.Difference(ss1)
	assert.True(t, s2.Equals(New().Add("was", "an", "interesting")))
}

func TestNegate2(t *testing.T) {
	ss1 := New().Add("a", "b", "c")
	ss2 := New().Add("c", "d", "e").Negate()
	pn := ss1.Intersection(ss2)
	assert.True(t, pn.Equals(New().Add("a", "b")))
}

func TestNegate1(t *testing.T) {
	ss1 := New().Add("a", "b", "c").Negate()
	ss2 := New().Add("c", "d", "e")
	pn := ss1.Intersection(ss2)
	assert.True(t, pn.Equals(New().Add("d", "e")))
}

func TestNegate12(t *testing.T) {
	ss1 := New().Add("a", "b", "c").Negate()
	ss2 := New().Add("c", "d", "e").Negate()
	pn := ss1.Intersection(ss2)
	assert.True(t, pn.Equals(New().Add("a", "b", "c", "d", "e")))
}

// This performs an equivalence test for two string slices
func checkEquivalence(t *testing.T, a []string, b []string) {
	// make a copy of a in c
	c := make([]string, len(a))
	copy(c, a)
	// look at every item in b and remove it from c
	for _, v := range b {
		found := false
		for i, v2 := range c {
			if v == v2 {
				// remove it from c by copying the last item over the found item and shortening the slice
				c[i], c = c[len(c)-1], c[:len(c)-1]
				found = true
				break
			}
		}
		assert.True(t, found, "Item '%v' not found in first list (%v)", v, c)
	}
	assert.Empty(t, c, "Extra items found in first list: '%v'", c)
}
