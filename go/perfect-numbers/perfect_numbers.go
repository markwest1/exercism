package perfect

const testVersion = 1

// Classification provides type-safety.
type Classification int

// Classification constants.
const (
	ClassificationDeficient Classification = iota
	ClassificationPerfect
	ClassificationAbundant
)

// ErrOnlyPositive is returned if input is non-positive.
var ErrOnlyPositive error

// Classify identifies the input number in one of three Nicomachus categories
// of perfect, abundant or deficient.
func Classify(n uint64) (class Classification, err error) {
	if n <= 0 {
		err = ErrOnlyPositive
		return
	}

	sum := uint64(0)

	for _, f := range factor(n) {
		sum += f
	}

	if sum == n {
		class = ClassificationPerfect
	} else if sum > n {
		class = ClassificationAbundant
	} else if sum < n {
		class = ClassificationDeficient
	}

	return
}

func factor(n uint64) (factors []uint64) {
	var f uint64
	for f = 1; f < n; f++ {
		if n%f == 0 {
			factors = append(factors, f)
		}
	}

	return
}
