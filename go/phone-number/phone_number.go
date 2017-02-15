package phonenumber

import (
	"bytes"
	"fmt"
	"unicode"
)

const testVersion = 1

// Number receives user input and returns a digit-only phone number
// or an error if the input does not conform to the following:
//   - The input must have exactly 10 or 11 digits
//   - If the input has 11 digits, the first digit must be a 1
func Number(in string) (string, error) {
	var buf bytes.Buffer

	digits := 0
	first := 'X'

	for _, r := range in {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)

			if digits == 0 {
				first = r
			}

			digits++
		}
	}

	if digits == 10 {
		return buf.String(), nil
	}

	if digits == 11 && first == '1' {
		return buf.String()[1:digits], nil
	}

	return "", fmt.Errorf("invalid input: %q", in)
}

// AreaCode returns the three-digit area code of a valid phone
// number; if the input number is invalid, an error is returned.
func AreaCode(in string) (string, error) {
	pn, err := Number(in)
	if err != nil {
		return "", err
	}

	return pn[0:3], nil
}

// Format receives an un-formatted phone number "1234567890" and
// returns "(123) 456-7890" if it is valid; otherwise an error
// is returned.
func Format(in string) (string, error) {
	pn, err := Number(in)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", pn[0:3], pn[3:6], pn[6:10]), nil
}
