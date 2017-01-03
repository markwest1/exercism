package wordcount

import (
	"bytes"
	"strings"
	"unicode"
)

const testVersion = 2

// Frequency is the return type.
type Frequency map[string]int

// WordCount returns a map of strings to the number of times the word occurs in
// the phrase.
func WordCount(phrase string) Frequency {
	words := strings.Split(phrase, " ")

	f := make(Frequency)
	for _, x := range words {
		if w := lettersAndDigitsOnly(x); w != "" {
			if _, ok := f[w]; !ok {
				f[w] = 1
			} else {
				f[w]++
			}
		}
	}

	return f
}

func lettersAndDigitsOnly(in string) string {
	var buf bytes.Buffer

	for _, r := range in {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			buf.WriteRune(unicode.ToLower(r))
		}
	}

	return buf.String()
}
