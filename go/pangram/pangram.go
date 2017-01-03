package pangram

import "unicode"

var (
	testVersion int = 1

	alphabet = map[rune]bool{
		rune('a'): true, rune('b'): true, rune('c'): true, rune('d'): true,
		rune('e'): true, rune('f'): true, rune('g'): true, rune('h'): true,
		rune('i'): true, rune('j'): true, rune('k'): true, rune('l'): true,
		rune('m'): true, rune('n'): true, rune('o'): true, rune('p'): true,
		rune('q'): true, rune('r'): true, rune('s'): true, rune('t'): true,
		rune('u'): true, rune('v'): true, rune('w'): true, rune('x'): true,
		rune('y'): true, rune('z'): true,
	}

	alphabetLength = len(alphabet)
)

// IsPangram returns true if string t contains every letter (case-insensitive)
// of the ASCII alphabet
func IsPangram(t string) bool {
	if len(t) < alphabetLength {
		return false
	}

	uniq := make(map[rune]bool)

	for _, r := range t {
		lc := unicode.ToLower(r)
		if _, ok := alphabet[lc]; !ok {
			continue
		}

		if _, ok := uniq[lc]; ok {
			continue
		}

		uniq[lc] = true
		if len(uniq) == alphabetLength {
			return true
		}
	}

	return len(uniq) == alphabetLength
}
