package isogram

import "unicode"

var testVersion = 1

// IsIsogram determines if the input string s is an isogram. An isogram (also
// known as a "nonpattern word") is a word or phrase without a repeating
// letter.
func IsIsogram(s string) bool {
	// A map is used to test for uniqueness. If a rune which is also a letter
	// already exists in the map, the rune is not unique to the string s.
	uniq := map[rune]bool{}

	for _, r := range s {
		if unicode.IsLetter(r) {
			lc := unicode.ToLower(r)
			if _, ok := uniq[lc]; ok {
				return false
			}

			uniq[lc] = true
		}
	}

	return true
}
