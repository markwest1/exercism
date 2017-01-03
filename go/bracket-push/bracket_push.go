package brackets

import "strings"

const testVersion = 4

var (
	parens     = "()"
	braces     = "{}"
	brackets   = "[]"
	openDelims = "({["
	allDelims  = parens + braces + brackets
)

// Bracket returns true if the order and placement of all opening and closing
// delimeters match.
func Bracket(s string) (match bool, err error) {
	// An empty expression is considered a match
	if s == "" {
		match = true
		return
	}

	opens := make([]string, 0)

	for i := 0; i < len(s); i++ {
		d := s[i : i+1]
		if !delimeter(d) {
			return
		}

		if open(d) {
			// Save open delimeters
			opens = append(opens, d)
		} else {
			if len(opens) == 0 {
				// No matching open delimeter
				return
			}

			length := len(opens) - 1
			lastOpen := opens[length]

			if sameType(d, lastOpen) {
				opens = opens[:length]
			} else {
				// Delimeter type mismatch
				return
			}
		}
	}

	// All delimeters match if opens is empty
	match = len(opens) == 0

	return
}

func open(s string) bool {
	return strings.Contains(openDelims, s)
}

func paren(s string) bool {
	return strings.Contains(parens, s)
}

func brace(s string) bool {
	return strings.Contains(braces, s)
}

func bracket(s string) bool {
	return strings.Contains(brackets, s)
}

func delimeter(s string) bool {
	return strings.Contains(allDelims, s)
}

func sameType(s1, s2 string) bool {
	return paren(s1) && paren(s2) || brace(s1) && brace(s2) || bracket(s1) && bracket(s2)
}
