package say

import (
	"bytes"
	"fmt"
)

var num = []uint64{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	30, 40, 50, 60, 70, 80, 90,
}

var w1 = []string{
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

var w2 = []string{
	"one", "two", "three", "four", "five",
	"six", "seven", "eight", "nine", "ten",
	"eleven", "twelve", "thirteen", "fourteen", "fifteen",
	"sixteen", "seventeen", "eighteen", "nineteen", "twenty",
	"thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
}

var w3 = []string{
	"thousand", "million", "billion", "trillion", "quadrillion", "quintillion",
}

// Say converts decimal numbers in the range to spoken English.
func Say(n uint64) string {
	var buf bytes.Buffer

	// n = 0
	if n == 0 {
		return "zero"
	}

	// n >= 1000
	if n >= 1000 {
		x := make([]uint64, 0)

		for n > 0 {
			x = append(x, n%1000)
			n /= 1000
		}

		for i := len(x) - 1; i >= 0; i-- {
			if x[i] == 0 {
				continue
			}

			if i >= 1 {
				nn := x[i]

				if nn >= 100 {
					spaceIfNotEmpty(&buf)
					buf.WriteString(lessThanOneThousand(nn))
					for nn >= 100 {
						nn -= 100
					}
				}

				if nn > 0 {
					spaceIfNotEmpty(&buf)
					buf.WriteString(lessThanOneHundred(nn))
				}

				buf.WriteString(fmt.Sprintf(" %s", w3[i-1]))
			}
		}

		n = x[0]
	}

	// 100 <= n < 1000
	if n < 1000 && n >= 100 {
		spaceIfNotEmpty(&buf)
		buf.WriteString(lessThanOneThousand(n))

		for n >= 100 {
			n -= 100
		}
	}

	// 0 < n < 100
	if n > 0 {
		spaceIfNotEmpty(&buf)
		buf.WriteString(lessThanOneHundred(n))
	}

	return buf.String()
}

func lessThanOneHundred(n uint64) string {
	if n >= 100 {
		return ""
	}

	var buf bytes.Buffer

	for i := 26; i >= 0; i-- {
		if n >= num[i] {
			if buf.Len() > 0 {
				if i >= 0 && i < 9 {
					buf.WriteString("-")
				} else {
					buf.WriteString(" ")
				}
			}
			buf.WriteString(w2[i])
			n -= num[i]
		}
	}

	return buf.String()
}

func lessThanOneThousand(n uint64) string {
	if n >= 1000 {
		return ""
	}

	q := n / 100

	return fmt.Sprintf("%s hundred", w1[q-1])
}

func spaceIfNotEmpty(buf *bytes.Buffer) {
	if buf.Len() > 0 {
		buf.WriteString(" ")
	}
}
