package foodchain

import (
	"bytes"
	"fmt"
)

var testVersion int = 3

var swallowed = []string{
	"fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse",
}

var first string = "I know an old lady who swallowed a %s."
var refrain = "She swallowed the %s to catch the %s"
var insideHer = " that wriggled and jiggled and tickled inside her."
var last string = "I don't know why she swallowed the fly. Perhaps she'll die."

var second = map[int]string{
	1: "It wriggled and jiggled and tickled inside her.",
	2: "How absurd to swallow a bird!",
	3: "Imagine that, to swallow a cat!",
	4: "What a hog, to swallow a dog!",
	5: "Just opened her throat and swallowed a goat!",
	6: "I don't know how she swallowed a cow!",
	7: "She's dead, of course!",
}

func Verse(v int) string {
	if v < 1 || v > 8 {
		return ""
	}

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf(first, swallowed[v-1]))

	if v > 1 {
		buf.WriteString(fmt.Sprintf("\n%s", second[v-1]))
	}

	if v < 8 {
		for i := v; i > 1; i-- {
			buf.WriteString(fmt.Sprintf("\n"+refrain, swallowed[i-1], swallowed[i-2]))

			if i != 3 {
				buf.WriteString(".")
			} else {
				buf.WriteString(insideHer)
			}
		}

		buf.WriteString(fmt.Sprintf("\n%s", last))
	}

	return buf.String()
}

func Verses(s, e int) string {
	if e < s || s < 1 || e > 8 {
		return ""
	}

	var buf bytes.Buffer

	for v := s; v <= e; v++ {
		buf.WriteString(fmt.Sprintf("%s", Verse(v)))

		if v < e {
			buf.WriteString("\n\n")
		}
	}

	return buf.String()
}

func Song() string {
	return Verses(1, 8)
}
