package queenattack

import (
	"fmt"
	"math"
	"strings"
)

var testVersion = 1

const (
	rows = "12345678"
	cols = "abcdefgh"
)

// CanQueenAttack positions two queens on a chess board and indicates whether
// or not they are positioned so that they can attack each other.
func CanQueenAttack(w, b string) (attack bool, err error) {
	if offBoard(w) || offBoard(b) {
		err = fmt.Errorf("invalid position")
		return
	}

	if w == b {
		err = fmt.Errorf("two pieces cannot occupy the same position")
		return
	}

	attack = sameColumn(w, b) || sameRow(w, b) || sameDiagonal(w, b)

	return
}

// offBoard returns true if the input position is invalid
func offBoard(p string) bool {
	if len(p) != 2 {
		return true
	}

	if !strings.ContainsAny(p[0:1], "abcdefgh") {
		return true
	}

	if !strings.ContainsAny(p[1:], "12345678") {
		return true
	}

	return false
}

// sameColumn returns true if two positions are in the same column
func sameColumn(w, b string) bool {
	return w[0:1] == b[0:1]
}

// sameRow returns true if tow positions are in the same row
func sameRow(w, b string) bool {
	return w[1:] == b[1:]
}

// sameDiagonal returns true if two positions are on the same diagonal
func sameDiagonal(w, b string) bool {
	// Calculate horizontal distance
	hd := strings.Index(cols, w[0:1]) - strings.Index(cols, b[0:1])

	// Calculate vertical distance
	vd := strings.Index(rows, w[1:]) - strings.Index(rows, b[1:])

	// Positiions are on the same diagonal if the horizontal and vertical
	// distances between the pieces are equal
	return math.Abs(float64(hd)) == math.Abs(float64(vd))
}
