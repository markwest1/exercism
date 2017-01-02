package triangle

import "math"

const testVersion = 3

type Kind int

// Type-safe enumerations in other programming languages translate to auto-
// incrementing constants of a type alias in Go
const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

// KindFromSides determines the kind of triangle given the lengths of its sides.
func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}

	if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		return NaT
	}

	if math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1) {
		return NaT
	}

	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	if a+b < c || b+c < a || a+c < b {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}

	return Sca
}
