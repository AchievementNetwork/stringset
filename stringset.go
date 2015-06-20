// Package stringset provides the capabilities of a standard Set with strings, including
// the standard set operations and a specialized ability to have a "negative" set
// for intersections.
//
// Negative sets are useful for inverting the sense of a set when combining them
// using Intersection. Negative sets do not change the behavior of any other
// set operations, including Union and Difference.  Negative operations were created
// for managing a set of tags, where it is useful to be able to search for items
// that contain some tags but do not contain others.
//
// The API is designed to be chainable so that common operations become one-liners,
// and it is designed to be easy to use with slices of strings.
//
// This package does not return errors. Set operations should be fast and
// chainable; adding items more than once, or attempting to remove things
// that don't exist are not errors.
package stringset

// A StringSet is implemented using a map[string]bool -- strings are copied
// when added to the set.
type StringSet struct {
	IsNegative bool
	content    map[string]bool
}

// New constructs an empty StringSet.
func New() *StringSet {
	ss := new(StringSet)
	ss.content = make(map[string]bool)
	return ss
}

// Negate flips the IsNegative bit.
func (ss *StringSet) Negate() *StringSet {
	ss.IsNegative = !ss.IsNegative
	return ss
}

// Add puts one or more strings into the set.
// If a string is already present the set is unchanged.
func (ss *StringSet) Add(sa ...string) *StringSet {
	for _, s := range sa {
		ss.content[s] = true
	}
	return ss
}

// Delete removes one or more items from a set if they exist in the set.
func (ss *StringSet) Delete(sa ...string) *StringSet {
	for _, s := range sa {
		delete(ss.content, s)
	}
	return ss
}

// Length returns the number of items currently in the set.
func (ss *StringSet) Length() int {
	return len(ss.content)
}

// Contains returns whether a given string is in the set.
func (ss *StringSet) Contains(s string) bool {
	_, ok := ss.content[s]
	return ok
}
