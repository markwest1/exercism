package stringset

import (
	"bytes"
	"fmt"
)

const testVersion = 4

// Set is a collection of unique string values
type Set map[string]bool

// New creates a new Set
func New() Set {
	return make(map[string]bool)
}

// NewFromSlice creates a new set populated with all unique strings in seed
func NewFromSlice(seed []string) Set {
	set := New()
	if seed != nil {
		for _, s := range seed {
			if !set.Has(s) {
				set[s] = true
			}
		}
	}

	return set
}

// String creates a string representation of a set
func (s Set) String() string {
	var buf bytes.Buffer
	buf.WriteString("{")

	count := 0
	for k := range s {
		buf.WriteString(fmt.Sprintf("%q", k))
		count++
		if count != len(s) {
			buf.WriteString(", ")
		}
	}

	buf.WriteString("}")
	return buf.String()
}

// IsEmpty returns true of set contains zero items, otherwise false
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// Has returns true if set contains element e, otherwise false
func (s Set) Has(e string) bool {
	_, has := s[e]
	return has
}

// Subset returns true if s1 is a subset of s2
func Subset(s1, s2 Set) bool {
	for e1 := range s1 {
		if !s2.Has(e1) {
			return false
		}
	}

	return true
}

// Disjoint returns true if s1 and s2 share no elements
func Disjoint(s1, s2 Set) bool {
	if len(s1) == 0 || len(s2) == 0 {
		return true
	}

	for e1 := range s1 {
		if s2.Has(e1) {
			return false
		}
	}

	return true
}

// Equal returns true if s1 and s2 are identical
func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}

	for e1 := range s1 {
		if !s2.Has(e1) {
			return false
		}
	}

	return true
}

// Add appends the string n to Set s if n is unique in s
func (s Set) Add(n string) {
	if s.IsEmpty() || !s.Has(n) {
		s[n] = true
	}
}

// Intersection returns a set containing all elements shared in Sets s1 and s2
func Intersection(s1, s2 Set) Set {
	if len(s1) == 0 {
		return s1
	}

	if len(s2) == 0 {
		return s2
	}

	for e1 := range s1 {
		if !s2.Has(e1) {
			delete(s1, e1)
		}
	}

	return s1
}

// Difference returns a set containing all non-shared elements of s1 and s2
func Difference(s1, s2 Set) Set {
	if s1.IsEmpty() || s2.IsEmpty() {
		return s1
	}

	for e1 := range s1 {
		if s2.Has(e1) {
			delete(s1, e1)
		}
	}

	return s1
}

// Union returns a Set containing all shared and non-shared members of Sets s1 and s2
func Union(s1, s2 Set) Set {
	if len(s1) == 0 {
		return s2
	}

	if len(s2) == 0 {
		return s1
	}

	for e2 := range s2 {
		s1.Add(e2)
	}

	return s1
}
