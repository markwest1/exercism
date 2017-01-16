package strand

import "bytes"

const testVersion = 3

var complements = map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}

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
