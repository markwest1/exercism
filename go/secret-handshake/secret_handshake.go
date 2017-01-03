package secret

var testVersion int = 1

var actionMap = map[uint]string{
	1: "wink",
	2: "double blink",
	4: "close your eyes",
	8: "jump",
}

func Handshake(code uint) (actions []string) {
	for m := uint(1); m < 16 && m <= code; m <<= 1 {
		if mask := code & m; mask != 0 {
			actions = append(actions, actionMap[m])
		}
	}

	if mask := code & 16; mask != 0 {
		actions = Reverse(actions)
	}

	return
}

func Reverse(actions []string) (reversed []string) {
	size := len(actions)
	if size == 0 {
		return
	}

	reversed = make([]string, size)

	for i := 0; i < size; i++ {
		reversed[size-i-1] = actions[i]
	}

	return
}
