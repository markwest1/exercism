package strain

// Ints is a collection of integers
type Ints []int

// Lists is a collection of a collection of integers
type Lists [][]int

// Strings is a collection of strings
type Strings []string

// Keep keeps all members of i that resolve to true in predicate.
func (i Ints) Keep(predicate func(int) bool) Ints {
	if i == nil {
		return nil
	}

	o := make(Ints, 0)

	for _, x := range i {
		if predicate(x) {
			o = append(o, x)
		}
	}

	return o
}

// Discard keeps all members of i that resolve to false in predicate.
func (i Ints) Discard(predicate func(int) bool) Ints {
	if i == nil {
		return nil
	}

	o := make(Ints, 0)

	for _, x := range i {
		if !predicate(x) {
			o = append(o, x)
		}
	}

	return o
}

// Keep keeps all lists that resolve to true in predicate.
func (l Lists) Keep(predicate func([]int) bool) Lists {
	if l == nil {
		return nil
	}

	o := make(Lists, 0)

	for _, x := range l {
		if predicate(x) {
			o = append(o, x)
		}
	}

	return o
}

// Keep keeps all members of s that resolve to true in predicate.
func (s Strings) Keep(predicate func(string) bool) Strings {
	if s == nil {
		return nil
	}

	o := make(Strings, 0)

	for _, w := range s {
		if predicate(w) {
			o = append(o, w)
		}
	}

	return o
}
