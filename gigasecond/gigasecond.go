package gigasecond

import "time"

const testVersion = 4

// AddGigasecond calculates the date and time on which a person becomes
// one billion seconds of age
func AddGigasecond(birthday time.Time) time.Time {
	return birthday.Add(time.Second * 1000000000)
}
