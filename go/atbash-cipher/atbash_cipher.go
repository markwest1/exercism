package atbash

import (
	"bytes"
	"unicode"
)

const testVersion = 1

var codec = map[rune]rune{
	'a': 'z', 'b': 'y', 'c': 'x', 'd': 'w', 'e': 'v', 'f': 'u', 'g': 't', 'h': 's',
	'i': 'r', 'j': 'q', 'k': 'p', 'l': 'o', 'm': 'n', 'n': 'm', 'o': 'l', 'p': 'k',
	'q': 'j', 'r': 'i', 's': 'h', 't': 'g', 'u': 'f', 'v': 'e', 'w': 'd', 'x': 'c',
	'y': 'b', 'z': 'a',
}

// Atbash implements the atbash cipher, an ancient encryption system
func Atbash(in string) string {
	var buf bytes.Buffer

	count := 0
	for _, r := range in {
		isLetter := unicode.IsLetter(r)

		if isLetter {
			r = codec[unicode.ToLower(r)]
		}

		append := isLetter || unicode.IsDigit(r)

		if append {
			insertSpace := count != 0 && count%5 == 0

			if insertSpace {
				buf.WriteRune(' ')
			}

			buf.WriteRune(r)
			count++
		}
	}

	return buf.String()
}
