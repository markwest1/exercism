package strand

import "bytes"

const testVersion = 3

var complements = map[rune]rune{
	rune('G'): rune('C'),
	rune('C'): rune('G'),
	rune('T'): rune('A'),
	rune('A'): rune('U'),
}

// ToRNA returns the RNA complement for a given DNA strand.
func ToRNA(strand string) string {
	var buf bytes.Buffer
	for _, r := range strand {
		if c, ok := complements[r]; ok {
			buf.WriteRune(c)
		}
	}

	return buf.String()
}
