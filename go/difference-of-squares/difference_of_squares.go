package diffsquares

// SquareOfSums calculates the sum of the first n natural numbers squared.
func SquareOfSums(n int) int {
	var sum int

	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum * sum
}

// SumOfSquares calculates the sum of squares of the first n natural numbers.
func SumOfSquares(n int) int {
	var sum int

	for i := 1; i <= n; i++ {
		sum += i * i
	}

	return sum
}

// Difference calculates the difference between SquareOfSums and SumOfSquares.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
