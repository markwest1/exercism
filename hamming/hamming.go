package hamming

import "fmt"

const testVersion = 5

// Distance calculates the hamming distance between two strands of DNA
func Distance(a, b string) (d int, err error) {
	if len(a) != len(b) {
		err = fmt.Errorf("strands must have equal length")
		return
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			d++
		}
	}

	return
}
