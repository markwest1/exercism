package prime

const testVersion = 2

// Factors returns the prime factors of in, in ascending order.
func Factors(q int64) []int64 {
	factors := make([]int64, 0)

	if q > 1 {
		f := int64(2)

		for q > 1 {
			r := q % f

			if r == 0 {
				factors = append(factors, f)
				q = q / f
			} else {
				f = next(f)
			}
		}
	}

	return factors
}

// Get the next prime factor
func next(last int64) int64 {
	// Other than 2 ...
	if last == 2 {
		return 3
	}

	// Even numbers are not prime
	return last + 2
}
