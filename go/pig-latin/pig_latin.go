package igpay

import (
	"bytes"
	"strings"
)

var (
	vowels = "aeiou"
	suffix = "ay"
)

// PigLatin translates an English word or list of space-separated words, in,
// into their pig-latin equivalents using the following rules:
//
// 1. If a word begins with a vowel sound, add "ay" to the end of the word.
//    Vowel sounds include "xr" as in "xray" and "yt" as in "yttria", but not
//    "x" as in "xenon" or "y" as in "yellow".
//
// 2. If a word begins with a consonant sound, move the letters for the
//    consonant sound concatenated with "ay" to the end of the word. Consonant
//    sounds include "qu" as in "queen" and "squ" as in "square".
func PigLatin(in string) string {
	if in == "" {
		return ""
	}

	var words []string
	if strings.Contains(in, " ") {
		words = strings.Split(in, " ")
	} else {
		words = make([]string, 1)
		words[0] = in
	}

	var translation bytes.Buffer
	for i, word := range words {
		var t string

		if vowelStart(word) {
			t = in + suffix
		} else if start, body := consonantStart(word); start != "" && body != "" {
			t = body + start + suffix
		}

		if i > 0 {
			t = " " + t
		}

		translation.WriteString(t)
	}

	return translation.String()
}

// vowelStart returns true if the string in starts with a vowel sound
func vowelStart(in string) bool {
	if len(in) == 0 {
		return false
	}

	// True if any of "aeiou" appear at the beginning of the word
	answer := strings.ContainsAny(in[0:1], vowels)

	// "xr" or "yt" at the beginning of a word are vowel sounds
	if !answer && len(in) > 1 {
		answer = in[0:2] == "yt" || in[0:2] == "xr"
	}

	return answer
}

// consonantStart returns the starting consonants and the rest of the word
// in two separate strings
func consonantStart(in string) (string, string) {
	if len(in) > 1 && in[0:2] == "qu" {
		return in[0:2], in[2:]
	}

	if len(in) > 2 && in[0:3] == "squ" {
		return in[0:3], in[3:]
	}

	index := 1
	for index < len(in) && !strings.ContainsAny(in[0:index], vowels) {
		index++
	}

	return in[0 : index-1], in[index-1:]
}
