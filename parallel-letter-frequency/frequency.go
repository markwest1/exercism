package letter

import "sync"

type FreqMap map[rune]int

// Frequency squentially counts the number of times a rune appears in a string
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency starts multiple goroutines to concurrently count the
// number of times a rune appears in a slice of strings
func ConcurrentFrequency(r []string) FreqMap {
	return concurrentFrequecyWithWaitGroup(r)
	// return concurrentFrequencyWithCtrlChannel(r)
}

// concurrentFrequecyWithWaitGroup uses a sync.WaitGroup to concatinate results
// of concurrent goroutines
func concurrentFrequecyWithWaitGroup(r []string) (fm FreqMap) {
	wg := sync.WaitGroup{}
	fms := make([]FreqMap, len(r))

	for i, s := range r {
		wg.Add(1)
		go func(i int, s string) {
			defer wg.Done()
			fms[i] = Frequency(s)
		}(i, s)
	}

	wg.Wait()

	fm = FreqMap{}
	for _, frMap := range fms {
		for k, v := range frMap {
			fm[k] += v
		}
	}

	return
}

// concurrentFrequencyWithCtrlChannel uses a "control" channel to concatinate
// results of concurrent goroutines
func concurrentFrequencyWithCtrlChannel(r []string) (fm FreqMap) {
	c := make(chan FreqMap)
	q := make(chan int)
	fm = FreqMap{}

	for _, s := range r {
		st := s
		go func() {
			c <- Frequency(st)
			q <- 0
		}()
	}

	rc := 0

	for {
		select {
		case fma := <-c:
			for k, v := range fma {
				fm[k] += v
			}
		case <-q:
			if rc++; rc == len(r) {
				return
			}
		}
	}
}
