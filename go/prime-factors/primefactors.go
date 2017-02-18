package prime

const testVersion = 2

// Factors returns the prime factors of q, in ascending order.
// This is an iteration on doozr's submission:
// http://exercism.io/submissions/ed668b42b10749528fea23296e32822e
func Factors(q int64) []int64 {
	factors := make([]int64, 0)

	// Two is the only even prime number
	for q%2 == 0 {
		q = q / 2
		factors = append(factors, 2)
	}

	// Even numbers are not prime
	for i := int64(3); i*i <= q; i += 2 {
		for q%i == 0 {
			q = q / i
			factors = append(factors, i)
		}
	}

	if q > 1 {
		factors = append(factors, q)
	}

	return factors
}
