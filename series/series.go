package slice

// All returns a list of all substrings of s with length n.
func All(n int, s string) []string {
	all := []string{}
	for i := 0; n+i <= len(s); i++ {
		all = append(all, s[i:n+i])
	}

	return all
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

// First includes boundary checks to safely implement UnsafeFirst
func First(n int, s string) (first string, ok bool) {
	ok = n <= len(s)
	if ok {
		first = UnsafeFirst(n, s)
	}

	return
}
