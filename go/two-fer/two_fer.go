// Package twofer returns one for X one for me
package twofer

// ShareWith returns the string "One for 'name', and one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}
