package bob

import (
	"strings"
	"unicode"
)

const (
	testVersion     = 2
	InquiryResponse = "Sure."
	YellResponse    = "Whoa, chill out!"
	SilenceResponse = "Fine. Be that way!"
	DefaultResponse = "Whatever."
)

// Hey returns Bob's response to an inquiry, yell, silence or babble.
func Hey(in string) (out string) {
	t := strings.Trim(in, "\n\r\t\v\u00a0\u2002 ")

	if t == "" {
		out = SilenceResponse
		return
	}

	var noLetters = true
	for _, r := range t {
		if unicode.IsLetter(r) {
			noLetters = false
			break
		}
	}

	if noLetters == false {
		var isYell bool = true
		for _, r := range t {
			if unicode.IsLetter(r) {
				if unicode.IsLower(r) {
					isYell = false
				}
			}
		}
		if isYell {
			out = YellResponse
			return
		}
	}

	runes := []rune(t)
	if runes[len(runes)-1] == rune('?') {
		out = InquiryResponse
		return
	}

	out = DefaultResponse

	return
}
