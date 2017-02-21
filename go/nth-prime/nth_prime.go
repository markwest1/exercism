package prime

const testVersion = 1

var primes []int

func init() {
	primes = sieve(1e6)
}

// Nth returns the n-th prime number, assuming the first is 2.
func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}

	return primes[n-1], true
}

// sieve finds prime numbers from 2 up to a given number using the Sieve of
// Eratosthenes algorithm.
func sieve(limit int) []int {
	prime := []int{}
	composite := map[int]bool{}

	// Loop through the integers [2, limit]
	for p := 2; p <= limit; p++ {
		// Skip composite numbers
		if _, marked := composite[p]; marked {
			continue
		}

		// Sifted out a prime number
		prime = append(prime, p)

		// Multiples of a prime number are composite
		i := 2
		prod := p * i
		for prod <= limit {
			if !composite[prod] {
				composite[prod] = true
			}

			i++
			prod = p * i
		}
	}

	// The sieve is complete, return the list of prime numbers
	return prime
}
