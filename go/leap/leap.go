package leap

const testVersion = 2

// IsLeapYear returns true if y is a leap year in the Gregorian calendar.
func IsLeapYear(y int) bool {
	isDivisibleBy4 := (y % 4) == 0
	isNotDivisibleBy100 := (y % 100) != 0
	isDivisibleBy400 := (y % 400) == 0

	return isDivisibleBy4 && (isNotDivisibleBy100 || isDivisibleBy400)
}
