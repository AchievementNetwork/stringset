package stringset

// Clone duplicates a StringSet.
func (ss *StringSet) Clone() *StringSet {
	clone := New()
	for k := range ss.content {
		clone.content[k] = true
	}
	clone.IsNegative = ss.IsNegative
	return clone
}

// This is a symmetric difference helper that intersects two sets and returns a new set.
// It helps with the negative operations by handling the case of
// the right hand argument possibly being negative.
//
// abc & cde == c
//
// abc & !cde == ab
func (ss1 *StringSet) intersection(ss2 *StringSet) *StringSet {
	// if we have 2 positive sets we can optimize for length
	if !ss2.IsNegative {
		l1 := len(ss1.content)
		l2 := len(ss2.content)
		if l2 < l1 {
			ss2, ss1 = ss1, ss2
		}
	}

	intersection := New()
	for k := range ss1.content {
		if _, ok := ss2.content[k]; ok != ss2.IsNegative {
			intersection.Add(k)
		}
	}
	return intersection
}

// Intersection returns the symmetric difference of the two sets,
// with the possibility that either or both sets may be negative.
// It's the equivalent of a boolean AND with either or both of the
// targets being inverted.
//
// This returns a new set -- both operands are left unchanged.
//
// abc & cde == c
//
// abc & !cde == ab
//
// !abc & cde == de
//
// !abc & !cde == !abcde
func (ss1 *StringSet) Intersection(ss2 *StringSet) *StringSet {
	var r *StringSet
	switch {
	case !ss1.IsNegative && !ss2.IsNegative:
		r = ss1.intersection(ss2)
	case !ss1.IsNegative && ss2.IsNegative:
		r = ss1.intersection(ss2)
	case ss1.IsNegative && !ss2.IsNegative:
		r = ss2.intersection(ss1)
	case ss1.IsNegative && ss2.IsNegative:
		r = ss1.Union(ss2).Negate()
	}
	return r
}

// Union generates the union (OR) of the two sets.
// Returns a new set and leaves the operands unchanged.
// IsNegative has no effect.
func (ss1 *StringSet) Union(ss2 *StringSet) *StringSet {
	union := New()
	for k := range ss1.content {
		union.Add(k)
	}
	for k := range ss2.content {
		union.Add(k)
	}
	return union
}

// Difference is an assymetric set difference.
// It subtracts the rhs from the lhs and returns a new set.
// IsNegative has no effect.
func (ss1 *StringSet) Difference(ss2 *StringSet) *StringSet {
	difference := New()
	for k := range ss1.content {
		if _, ok := ss2.content[k]; !ok {
			difference.Add(k)
		}
	}
	return difference
}

// Equals checks whether two string sets have the same members.
// IsNegative has no effect.
func (ss1 *StringSet) Equals(ss2 *StringSet) bool {
	if ss1.Length() != ss2.Length() {
		return false
	}
	// if two sets have the same length then we know that if there is a difference,
	// it must be the case that there are the same number of unique members
	// in each set. So we can iterate over one set, and know that if every
	// item was found, the two sets are equivalent
	for k := range ss1.content {
		if _, ok := ss2.content[k]; !ok {
			return false
		}
	}
	return true
}
