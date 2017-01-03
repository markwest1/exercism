package binary

import "fmt"

const (
	testVersion = 2

	zero = rune('0')
	one  = rune('1')
)

// ParseBinary returns the decimal equivalent of b, a binary number string.
func ParseBinary(b string) (v int, err error) {
	l := len(b)
	for i, r := range b {
		// Binary allows only two digits
		if r != one && r != zero {
			v = 0
			err = fmt.Errorf("invalid input: %q", b)
			return
		}

		if r == one {
			if i == len(b)-1 {
				v++
			} else {
				v += (2 << uint(l-i-2))
			}
		}
	}

	return
}
