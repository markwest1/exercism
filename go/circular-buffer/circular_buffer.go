package circular

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
//   type Buffer
//   func NewBuffer(size int) *Buffer
//   func (*Buffer) ReadByte() (byte, error)
//   func (*Buffer) WriteByte(c byte) error
//   func (*Buffer) Overwrite(c byte)
//   func (*Buffer) Reset() // put buffer in an empty state
//
// We chose the above API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Buffer is a circular buffer supporting both overflow-checked writes and
// unconditional, possibly overwriting, writes.
type Buffer []int

// NewBuffer creates a new buffer.
func NewBuffer(size int) *Buffer {
	b := make(Buffer, size)
	return &b
}

// ReadByte ...
func (b *Buffer) ReadByte() (c byte, e error) {
	return
}

// WriteByte ...
func (b *Buffer) WriteByte(c byte) (e error) {
	return
}

// Overwrite ...
func (b *Buffer) Overwrite(c byte) {
	return
}

// Reset ...
func (b *Buffer) Reset() {
	return
}
