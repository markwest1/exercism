package lsproduct

import (
	"fmt"
	"strconv"
)

const testVersion = 4

// LargestSeriesProduct calculates the largest product for a contiguous
// substring of digits of length span.
func LargestSeriesProduct(digits string, span int) (p int, err error) {
	// Respond to special conditions
	if span > len(digits) {
		err = fmt.Errorf("Span (%d) is greater than the number of digits (%d).",
			span, len(digits))
		return
	} else if span < 0 {
		err = fmt.Errorf("Span cannot be negative")
		return
	} else if span == 0 {
		p = 1
		return
	}

	// Transform the string digits to a slice of integers
	p = -1
	numbers := []int{}
	for _, r := range digits {
		n, e := strconv.Atoi(string(r))
		if e != nil {
			err = fmt.Errorf("\"%s\" contains non-digits", digits)
			return
		}

		numbers = append(numbers, n)
	}

	// Find the largest series product
	p = 0
	for i := 0; i <= len(numbers)-span; i++ {
		prod := product(numbers[i : i+span])
		if prod > p {
			p = prod
		}
	}

	return
}

// product calculates the product of a slice of integers.
func product(nums []int) (p int) {
	p = 1

	for _, n := range nums {
		if n == 0 {
			return 0
		}

		p *= n
	}

	return
}
