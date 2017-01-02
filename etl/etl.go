package etl

import "strings"

// Transform changes data format from old to new
func Transform(old map[int][]string) (new map[string]int) {
	new = make(map[string]int)
	for points, letters := range old {
		for _, letter := range letters {
			new[strings.ToLower(letter)] = points
		}
	}

	return
}
