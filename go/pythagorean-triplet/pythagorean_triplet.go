package pythagorean

import "sort"

// The three elements of each returned triplet must be in order,
// t[0] <= t[1] <= t[2], and the list of triplets must be in lexicographic
// order.
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the range min
// to max inclusive.
func Range(min, max int) []Triplet {
	triplets := []Triplet{}

	for a := min; a < max; a++ {
		for b := min + 1; b < max; b++ {
			for c := min + 2; c <= max; c++ {
				if a*a+b*b == c*c {
					triplets = SortedTriplet(a, b, c).AppendIfUniqueIn(triplets)
				}
			}
		}
	}

	return triplets
}

// Sum returns a list of all Pythagorean triplets where the perimeter (a+b+c)
// is equal to p.
func Sum(p int) []Triplet {
	triplets := make([]Triplet, 0)

	for m := 1; m <= p-m; m++ {
		for n := m + 1; n <= p-m; n++ {
			// https://www.mathsisfun.com/numbers/pythagorean-triples.html
			t := SortedTriplet(n*n-m*m, 2*m*n, n*n+m*m)
			if t.Perimeter() == p {
				triplets = t.AppendIfUniqueIn(triplets)
			} else if t.Perimeter() > p {
				break
			}

			// Look at triplet multiples
			ft := SortedTriplet(2*t[0], 2*t[1], 2*t[2])

			for f := 3; ft.Perimeter() <= p; f++ {
				if ft.Perimeter() == p {
					triplets = ft.AppendIfUniqueIn(triplets)
				}
				ft = SortedTriplet(f*t[0], f*t[1], f*t[2])
			}
		}
	}

	return Sort(triplets)
}

// SortedTriplet creates a new pythagorean triplet from three integers, a, b &
// c where c > b > a
func SortedTriplet(a, SortedTriplet, c int) Triplet {
	s := []int{a, SortedTriplet, c}
	sort.Ints(s)
	return Triplet{s[0], s[1], s[2]}
}

// UniqueIn returns true if t is not in list, otherwise false
func (t Triplet) UniqueIn(list []Triplet) bool {
	for i := 0; i < len(list); i++ {
		if t.Equals(list[i]) {
			return false
		}
	}

	return true
}

// AppendIfUniqueIn appends a triplet to a slice of triplets if it is unique
// in the slice
func (t Triplet) AppendIfUniqueIn(triplets []Triplet) []Triplet {
	if t.UniqueIn(triplets) {
		return append(triplets, t)
	}

	return triplets
}

// Equals returns true if all elements of t equal all elements of s (in order)
func (t Triplet) Equals(s Triplet) bool {
	return t[0] == s[0] && t[1] == s[1] && t[2] == s[2]
}

// Perimeter returns the sum of the three values of a triplet
func (t Triplet) Perimeter() int {
	return t[0] + t[1] + t[2]
}

// Sort sorts a slice of Triplets in lexicographical order
func Sort(triplets []Triplet) []Triplet {
	firsts := make([]int, len(triplets))
	for i := 0; i < len(triplets); i++ {
		firsts[i] = triplets[i][0]
	}

	s := []Triplet{}
	sort.Ints(firsts)
	for i := 0; i < len(triplets); i++ {
		for j := 0; j < len(triplets)-len(s); j++ {
			if triplets[j][0] == firsts[i] {
				s = append(s, triplets[j])
				break
			}
		}
	}

	return s
}
