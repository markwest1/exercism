package prime

const testVersion = 1

// Nth returns the n-th prime number, assuming the first is 2.
func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}

	return -1, false
}
