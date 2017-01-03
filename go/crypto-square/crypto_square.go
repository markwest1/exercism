package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

var testVersion = 2

// Encode returns a square-encoded cypher of text.
func Encode(text string) string {
	return square(normalize(text))
}

// normalize transforms input text by removing spaces and punctuation and
// changing upper-case letters to lower-case letters.
func normalize(text string) (normalized string) {
	for _, r := range text {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			normalized += string(unicode.ToLower(r))
		}
	}
	return
}

// square re-organizes normalized text into a rectange of size rows x columns
// where columns >= rows and columns - rows <= 1
func square(text string) string {
	length := len(text)
	if length == 0 {
		return ""
	}

	rows, columns := dimensions(length)

	chunks := make([]string, columns)
	for col := 0; col < columns; col++ {
		for row := 0; row < rows; row++ {
			if index := col + row*columns; index < length {
				if ru, ok := runeAt(text, index); ok {
					chunks[col] += string(ru)
				}
			} else {
				break
			}
		}
	}

	return strings.Join(chunks, " ")
}

// dimensions identifies the number of rows and columns required to produce a
// square cypher for normallized text of length l.
func dimensions(l int) (rows, columns int) {
	rows = int(math.Sqrt(float64(l)))
	columns = rows

	for rows*columns < l {
		if columns > rows {
			rows++
		} else {
			columns++
		}
	}

	return
}

// runeAt returns the rune at index of text
func runeAt(text string, index int) (r rune, ok bool) {
	if index >= len(text) {
		return
	}

	ok = true
	for i, ru := range text {
		if i == index {
			r = ru
			break
		}
	}

	return
}
