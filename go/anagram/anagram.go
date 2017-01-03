package anagram

import "strings"

// Detect identifies anagrams of subject from a list of candidates.
func Detect(subject string, candidates []string) (anagrams []string) {
	// Channels on which goroutines communicate information; one for indexes of
	// anagrams, one to indicate completion.
	ic := make(chan int)
	fc := make(chan bool)

	// Map each rune in the subject to the number of occurences in subject.
	sm := make(map[rune]int)
	for _, r := range strings.ToLower(subject) {
		if _, ok := sm[r]; !ok {
			sm[r] = 1
		} else {
			sm[r]++
		}
	}

	// Remember the number of candidates.
	numCandidates := len(candidates)

	// Start goroutines to detect anagrams.
	for i, c := range candidates {
		// Do not start goroutines for identical or diferent-length candidates.
		if strings.ToLower(c) == strings.ToLower(subject) || len(c) != len(subject) {
			numCandidates--
			continue
		}

		go anagram(sm, strings.ToLower(c), i, ic, fc)
	}

	// If zero goroutines started, detection algorithm is commplete.
	if numCandidates == 0 {
		return
	}

	// Listen for feedback from goroutines started above.
	count := 0
	for {
		select {
		case ai := <-ic:
			anagrams = append(anagrams, strings.ToLower(candidates[ai]))
		case <-fc:
			count++
			if count == numCandidates {
				return
			}
		}
	}
}

// anagram puts index into indexChan if candidate is an anagram of subject
func anagram(sm map[rune]int, candidate string, index int, indexChan chan int, finChan chan bool) {
	cm := make(map[rune]int)
	for _, r := range candidate {
		if _, ok := cm[r]; !ok {
			cm[r] = 1
		} else {
			cm[r]++
			if cm[r] > sm[r] {
				finChan <- true
				return
			}
		}
	}

	if same(sm, cm) {
		indexChan <- index
	}

	finChan <- true
}

// same returns true if sMap and cMap are the same length and contain identical
// key-value pairs.
func same(sMap, cMap map[rune]int) bool {
	for sr, sCount := range sMap {
		if cCount, ok := cMap[sr]; !ok {
			return false
		} else if cCount != sCount {
			return false
		}
	}

	return true
}
