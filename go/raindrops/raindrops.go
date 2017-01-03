package raindrops

import "fmt"

const testVersion = 2

// Convert transforms a number's representation based on its factors.
func Convert(n int) (r string) {
	if n%3 == 0 {
		r = "Pling"
	}

	if n%5 == 0 {
		r += "Plang"
	}

	if n%7 == 0 {
		r += "Plong"
	}

	if r == "" {
		r = fmt.Sprintf("%d", n)
	}

	return r
}
