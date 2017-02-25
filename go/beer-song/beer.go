package beer

import (
	"fmt"
	"log"
)

const testVersion = 1

// Verse returns a verse of the beer song.
func Verse(num int) (string, error) {
	return "", fmt.Errorf("not implemented")
}

// Verses returns verses upper to lower of the beer song.
func Verses(upper, lower int) (string, error) {
	return "", fmt.Errorf("not implemented")
}

// Song returns all verses of the beer song.
func Song() string {
	song, err := Verses(99, 1)
	if err != nil {
		log.Fatal(err)
	}

	return song
}
