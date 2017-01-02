package palindrome

import (
	"fmt"
	"math"
	"strconv"
)

// Direction provides a typesafe constant for finding greatest or least palindrome
type Direction int

const (
	testVersion = 1

	up Direction = iota
	down
)

// Product is the output structure
type Product struct {
	// palindromic, of course
	Product int
	// list of all possible two-factor factorizations of Product, within given limits, in order
	Factorizations [][2]int
}

// Products finds the min and max palindromic products of factors between fmin and fmax.
func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	// fmin must be less than fmax
	if fmin > fmax {
		err = fmt.Errorf("fmin > fmax")
		return
	}

	done := make(chan bool)
	minc := findPalindrome(fmin, fmax, up, done)
	maxc := findPalindrome(fmin, fmax, down, done)
	count := 0

	for {
		select {
		case pmin = <-minc:
		case pmax = <-maxc:
		case <-done:
			count++
			if count == 2 {
				if pmin.Product == 0 && pmax.Product == 0 {
					err = fmt.Errorf("No palindromes")
				}
				return
			}
		}
	}
}

// findPalindrome finds the first palindrome and all its factors
func findPalindrome(f1, f2 int, d Direction, done chan<- bool) (pc chan Product) {
	pc = make(chan Product)

	// Start a goroutine here
	go func() {
		var prod int
		var fact [][2]int

		if d == up {
			prod = math.MaxInt32
			for i := f1; i <= f2; i++ {
				for j := i; j <= f2; j++ {
					p := i * j
					if p < prod && isPalindrome(p) {
						prod = p
						fact = make([][2]int, 1)
						fact[0] = [2]int{i, j}
					} else if p == prod {
						fact = append(fact, [2]int{i, j})
					}
				}
			}
		} else if d == down {
			prod = math.MinInt32
			for i := f2; i >= f1; i-- {
				for j := i; j >= f1; j-- {
					p := i * j
					if p > prod && isPalindrome(p) {
						prod = p
						fact = make([][2]int, 1)
						fact[0] = [2]int{j, i}
					} else if p == prod {
						fact = append(fact, [2]int{j, i})
					}
				}
			}
		}

		if prod == math.MinInt32 || prod == math.MaxInt32 {
			prod = 0
		}

		pc <- Product{prod, fact}
		done <- true
	}()

	return pc
}

// isPalindrome determines if the number n is a palindrome
func isPalindrome(n int) bool {
	s := strconv.Itoa(n)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
