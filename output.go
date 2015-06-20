package stringset

import "strings"

// Strings returns all the items in the set as a slice of strings.
func (ss *StringSet) Strings() []string {
	a := make([]string, 0, len(ss.content))
	for s := range ss.content {
		a = append(a, s)
	}
	return a
}

// Join returns all the items in the set as a single string, in
// an arbitrarily-ordered list, separated by a separator.
func (ss *StringSet) Join(sep string) string {
	return strings.Join(ss.Strings(), sep)
}

// WrappedJoin does the same as Join but allows a prefix and a suffix.
func (ss *StringSet) WrappedJoin(prefix string, sep string, suffix string) string {
	return prefix + strings.Join(ss.Strings(), sep) + suffix
}
