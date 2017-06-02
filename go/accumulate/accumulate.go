package accumulate

const testVersion = 1

// Accumulate is like a Map
func Accumulate(given []string, converter func(string) string) []string {
	out := []string{}

	for _, s := range given {
		out = append(out, converter(s))
	}

	return out
}
