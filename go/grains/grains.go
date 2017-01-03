package grains

import "fmt"

const testVersion = 1

// Square computes the number of grains in square n of the chessboard
func Square(n int) (v uint64, err error) {
	if n <= 0 || n > 64 {
		err = fmt.Errorf("input must be an integer between 1 and 64")
	}

	v = 1
	if n == 1 {
		return
	}

	for i := 1; i < n; i++ {
		v *= 2
	}

	return
}

// Total computes the number of grains on the entire chessboard
func Total() (t uint64) {
	for i := 1; i <= 64; i++ {
		s, _ := Square(i)
		t += s
	}

	return
}
