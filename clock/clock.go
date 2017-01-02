package clock

import "fmt"

const testVersion = 4

type Clock struct {
	h int
	m int
}

// New creates a new clock
func New(hour, minute int) Clock {
	for minute > 59 {
		minute -= 60
		hour++
	}
	for minute < 0 {
		minute += 60
		hour--
	}

	for hour > 23 {
		hour -= 24
	}
	for hour < 0 {
		hour += 24
	}

	return Clock{hour, minute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func (c Clock) Add(minutes int) Clock {
	return New(c.h, c.m+minutes)
}
