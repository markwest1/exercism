package diamond

import (
	"bytes"
	"fmt"
)

const testVersion = 1

var letters = []rune{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
	'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

// Gen generates an alphabetical diamond with the rune c in the middle row.
func Gen(c byte) (string, error) {
	// Find the index of the input byte
	index, err := index(rune(c))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	// Length of a side of the square
	length := index*2 + 1

	// Create the diamond from top to bottom
	for row := 0; row < length; row++ {
		top := row <= index

		for col := 0; col < length; col++ {
			// Space is default
			r := ' '

			// In the top half of the square
			if top && col == index-row || col == length-(index-row+1) {
				r = letters[row]
			}

			// In the bottom half of the square
			if !top && col == -(index-row) || col == length+(index-row-1) {
				r = letters[length-row-1]
			}

			buf.WriteRune(r)
		}

		// Add line-ending
		buf.WriteString("\n")
	}

	return buf.String(), nil
}

func index(letter rune) (int, error) {
	for i := 0; i < len(letters); i++ {
		if letter == letters[i] {
			return i, nil
		}
	}

	return 0, fmt.Errorf("%q is not an upper-case letter", letter)
}
