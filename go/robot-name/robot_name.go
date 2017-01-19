package robotname

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"
)

var letters = []rune{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K',
	'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V',
	'W', 'X', 'Y', 'Z',
}

// Robot represents a robot.
type Robot struct {
	name string
}

// Init initializes the seed of rand.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Name returns the name of robot r; created if necessary.
func (r *Robot) Name() string {
	if r.name == "" {
		var buf bytes.Buffer

		buf.WriteRune(letters[rand.Intn(26)])
		buf.WriteRune(letters[rand.Intn(26)])
		buf.WriteString(strconv.Itoa(rand.Intn(900) + 100))

		r.name = buf.String()
	}

	return r.name
}

// Reset deletes the name of robot r.
func (r *Robot) Reset() {
	r.name = ""
}
