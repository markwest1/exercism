package transpose

import "bytes"

// Transpose changes the orientation of a slice of strings from horizontal to
// vertical (i.e. after transpose, read from top-to-bottom rather than left-to-
// right).
func Transpose(h []string) (v []string) {
	v = make([]string, 0)
	if h == nil || len(h) == 0 {
		return
	}

	// The vertical width (columns) and height (rows)
	vW := len(h)
	vH := maxLen(h)

	// To avoid extraneous memory allocation, use byte buffers rather than
	// string concatenation
	vb := make([]bytes.Buffer, 0)
	for i := 0; i < vH; i++ {
		vb = append(vb, bytes.Buffer{})
	}

	// Transpose, avoiding hanging spaces in the vertically-oriented rows
	for vCol := 0; vCol < vW; vCol++ {
		for vRow := 0; vRow < vH; vRow++ {
			if len(h[vCol]) > vRow {
				vb[vRow].WriteByte(h[vCol][vRow])
			} else if vRow < maxLen(h[vCol:]) {
				vb[vRow].WriteRune(' ')
			}
		}
	}

	// Transform byte buffer contents to the expected return value
	for i := 0; i < vH; i++ {
		v = append(v, vb[i].String())
	}

	return
}

// maxLen returns the maximum length of the rows in rr
func maxLen(rr []string) int {
	max := 0
	for i := 0; i < len(rr); i++ {
		if len(rr[i]) > max {
			max = len(rr[i])
		}
	}

	return max
}
