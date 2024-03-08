package bytestream

type Reader interface {
	Peek() string          // Peek at the next bytes in the buffer
	Pop(len uint64) error  // Remove `len` bytes from the buffer
	IsFinished() bool      // Is the stream finished (closed and fully popped)?
	BytesBuffered() uint64 // Number of bytes currently buffered (pushed and not popped)
	BytesPopped() uint64   // Total number of bytes cumulatively popped from stream
}

type Writer interface {
	Push(data string) error    // Push data to stream, but only as much as available capacity allows.
	Close() error              // Signal that the stream has reached its ending. Nothing more will be written.
	IsClosed() bool            // Has the stream been closed?
	AvailableCapacity() uint64 // How many bytes can be pushed to the stream right now?
	BytesPushed() uint64       // Total number of bytes cumulatively pushed to the stream
}

type ByteStream struct {
	capacity     uint64
	streamBuffer string
	closed       bool
	nRead        uint64
	nWrite       uint64
}

func (bs *ByteStream) IsFull() bool {
	return false
}

func (bs *ByteStream) IsEmpty() bool {
	return false
}

func (bs *ByteStream) Capacity() uint64 {
	return bs.capacity
}

// ==================================
// Implementation of Writer interface
// ==================================

func (bs *ByteStream) Push(data string) error {

	return nil
}

func (bs *ByteStream) Close() error {
	return nil
}

func (bs *ByteStream) IsClosed() bool {
	return true
}

func (bs *ByteStream) AvailableCapacity() uint64 {
	return 0
}

func (bs *ByteStream) BytesPushed() uint64 {
	return 0
}

// ==================================
// Implementation of Reader interface
// ==================================

func (bs *ByteStream) Peek() string {
	return ""
}

func (bs *ByteStream) Pop(len uint64) error {
	return nil
}

func (bs *ByteStream) IsFinished() bool {
	return false
}

func (bs *ByteStream) BytesBuffered() uint64 {
	return 0
}

func (bs *ByteStream) BytesPopped() uint64 {
	return 0
}
