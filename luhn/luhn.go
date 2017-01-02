package luhn

import (
	"strconv"
	"unicode"
)

var testVersion = 1

// Valid determines if n is valid per the Luhn formula. The Luhn formula is a
// simple checksum formula used to validate a variety of identification numbers,
// such as credit card numbers and Canadian Social Insurance Numbers.
func Valid(n string) bool {
	// Extract digit runes from n.
	list := digits(n)

	// If n has no digits, it cannot be valid.
	if len(list) == 0 {
		return false
	}

	// If the checksum is 0, n is valid per the Luhn formula.
	return checksum(list) == 0
}

// AddCheck adds a check digit to make raw valid per the Luhn formula and
// returns the original number plus that digit.
func AddCheck(raw string) string {
	i := 0
	n := raw + strconv.Itoa(i)

	for !Valid(n) && i < 10 {
		i++
		n = raw + strconv.Itoa(i)
	}

	return n
}

// digits extracts digit runes from a string n.
func digits(n string) (runes []rune) {
	// Extract digit runes from n
	for _, r := range n {
		// Ignore non-digits
		if !unicode.IsDigit(r) {
			continue
		}

		runes = append(runes, r)
	}

	return
}

// checksum calculates the Luhn checksum of digits.
func checksum(digits []rune) int {
	sum := 0
	count := 1

	// Start at the right-most digit of n.
	for i := len(digits) - 1; i >= 0; i-- {
		v, err := strconv.Atoi(string(digits[i]))

		// Ignore un-convertable runes
		if err != nil {
			continue
		}

		// Double the value of every second digit
		if count%2 == 0 {
			v *= 2
			// Subtract 9 if v doubled is 10 or more.
			if v >= 10 {
				v -= 9
			}
		}

		count++
		sum += v
	}

	return sum % 10
}
