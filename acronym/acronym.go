package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 1

// abbreviate returns an acronym for the verbose version of the string
func abbreviate(in string) (out string) {
	for _, f := range strings.FieldsFunc(in, isFieldDelimiter) {
		runes := []rune(f)
		out += string(unicode.ToUpper(runes[0]))

		for i := 1; i < len(runes); i++ {
			if unicode.IsLetter(runes[i-1]) && unicode.IsLower(runes[i-1]) {
				if unicode.IsLetter(runes[i]) && unicode.IsUpper(runes[i]) {
					out += string(runes[i])
				}
			}
		}
	}

	return
}

func isFieldDelimiter(r rune) bool {
	return unicode.IsSpace(r) || r == rune('-')
}
