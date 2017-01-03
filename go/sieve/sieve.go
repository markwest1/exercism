package sieve

const testVersion = 1

// Sieve finds prime numbers from 2 up to a given number using the Sieve of
// Eratosthenes algorithm.
func Sieve(limit int) []int {
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
