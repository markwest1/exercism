package circular

import "fmt"

const testVersion = 4

// Buffer is a circular buffer supporting both overflow-checked writes and
// unconditional, possibly overwriting writes.
type Buffer struct {
	data   []byte
	size   int
	oldest int
	latest int
	empty  bool
	full   bool
}

// NewBuffer creates a new buffer with size s.
func NewBuffer(s int) *Buffer {
	if s <= 0 {
		fmt.Printf("buffer size must be positive")
		return nil
	}

	return &Buffer{make([]byte, s), s, 0, 0, true, false}
}

// ReadByte returns the byte at index oldest and increments oldest. Returns an
// error if the buffer is empty.
func (b *Buffer) ReadByte() (c byte, e error) {
	if b.empty {
		e = fmt.Errorf("can't read from empty buffer")
		return
	}

	c = b.data[b.oldest]
	b.full = false
	b.oldest = (b.oldest + 1) % b.size
	b.empty = b.oldest == b.latest

	return
}

// WriteByte saves byte c at index latest and increments latest. Returns an
// error if the buffer is full.
func (b *Buffer) WriteByte(c byte) (e error) {
	if b.full {
		e = fmt.Errorf("can't write to full buffer")
		return
	}

	b.data[b.latest] = c
	b.empty = false
	b.latest = (b.latest + 1) % b.size
	b.full = b.latest == b.oldest

	return
}

// Overwrite writes byte c at index oldest and increments both oldest and latest
// if the buffer is full. If the buffer is not full, WriteByte(c) is called.
func (b *Buffer) Overwrite(c byte) {
	if b.full {
		b.data[b.oldest] = c
		b.latest = (b.latest + 1) % b.size
		b.oldest = (b.oldest + 1) % b.size
	} else {
		b.WriteByte(c)
	}

	return
}

// Reset resets the state of circular buffer to empty.
func (b *Buffer) Reset() {
	b.oldest = b.latest
	b.empty = true
	b.full = false
}
