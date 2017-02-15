package phonenumber

import "bytes"

const testVersion = 1

// Number receives user input and returns a digit-only phone number
// or an error if the input does not conform to the following:
//   - The input must have exactly 10 or 11 digits
//   - If the input has 11 digits, the first digit must be a 1
func Number(in string) (string, error) {
	var buf bytes.Buffer

	return buf.String(), nil
}

// AreaCode returns the three-digit area code of a valid phone
// number; if the input number is invalid, an error is returned.
func AreaCode(in string) (string, error) {
	var buf bytes.Buffer

	return buf.String(), nil
}

// Format receives an un-formatted input phone number and returns
// it as "(123) 456-7890" if it is valid; otherwise an error is
// returned.
func Format(in string) (string, error) {
	var buf bytes.Buffer

	return buf.String(), nil
}
