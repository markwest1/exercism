package summultiples

// MultipleSummer returns a function that, given a number, can find the sum of
// all the multiples of particular numbers up to but not including that number.
func MultipleSummer(divisors ...int) func(t int) int {
	return func(limit int) (sum int) {
		list := divisors

		if limit <= 0 || len(list) == 0 {
			return
		}

		multiplier := 0
		uniq := make(map[int]bool)

		for {
			multiplier++
			overLimitCount := 0

			for i := 0; i < len(list); i++ {
				p := list[i] * multiplier
				_, alreadyAdded := uniq[p]

				if !alreadyAdded && p < limit {
					uniq[p] = true
					sum += p
				}

				if p >= limit {
					overLimitCount++
				}
			}

			if overLimitCount == len(list) {
				break
			}
		}

		return
	}
}
