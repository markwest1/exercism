package twelve

import "bytes"

var (
	testVersion = 1

	gift = map[int]string{
		1:  "a Partridge in a Pear Tree.",
		2:  "two Turtle Doves",
		3:  "three French Hens",
		4:  "four Calling Birds",
		5:  "five Gold Rings",
		6:  "six Geese-a-Laying",
		7:  "seven Swans-a-Swimming",
		8:  "eight Maids-a-Milking",
		9:  "nine Ladies Dancing",
		10: "ten Lords-a-Leaping",
		11: "eleven Pipers Piping",
		12: "twelve Drummers Drumming",
	}

	day = map[int]string{
		1:  "first",
		2:  "second",
		3:  "third",
		4:  "fourth",
		5:  "fifth",
		6:  "sixth",
		7:  "seventh",
		8:  "eighth",
		9:  "ninth",
		10: "tenth",
		11: "eleventh",
		12: "twelfth",
	}

	onThe      = "On the "
	myTrueLove = " my true love gave to me, "
)

// Song returns a string containing the lyrics to all twelve verses of the song
// "The Twelve Days of Christmas"
func Song() string {
	buf := bytes.Buffer{}
	for i := 1; i <= 12; i++ {
		buf.WriteString(Verse(i) + "\n")
	}

	return buf.String()
}

// Verse returns a string containing the lyrics to the indexed verse of the
// song "The Twelve Days of Christmas"
func Verse(index int) string {
	buf := bytes.Buffer{}

	if index == 1 {
		return onThe + day[index] + myTrueLove + gift[index]
	}

	buf.WriteString(onThe + day[index] + myTrueLove)

	for i := index; i >= 2; i-- {
		buf.WriteString(gift[i] + ", ")
	}

	buf.WriteString("and " + gift[1])

	return buf.String()
}
