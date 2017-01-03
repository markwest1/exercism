package romannumerals

import (
	"bytes"
	"fmt"
)

const testVersion = 3

var roman = []string{
	"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I",
}

var arabic = []int{
	1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1,
}

// ToRomanNumeral converts the arabic integer in to the roman numeral out.
func ToRomanNumeral(in int) (out string, err error) {
	if in > 3999 || in < 1 {
		err = fmt.Errorf("invalid input: %q", in)
		return
	}

	var buf bytes.Buffer

	for i, a := range arabic {
		for in >= a && in > 0 {
			if in-a >= 0 {
				in -= a
				buf.WriteString(roman[i])
			}
		}
	}

	out = buf.String()
	return
}
