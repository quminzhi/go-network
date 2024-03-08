package bytestream

import "testing"

func TestByteStreamPush(t *testing.T) {
	bs := &ByteStream{capacity: 10, streamBuffer: "", closed: false, nRead: 0, nWrite: 0}

	err := bs.Push("abc")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Add more test cases for Push if needed
}

func TestByteStreamClose(t *testing.T) {
	bs := &ByteStream{capacity: 10, streamBuffer: "", closed: false, nRead: 0, nWrite: 0}

	err := bs.Close()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Add more test cases for Close if needed
}

func TestWriter(t *testing.T) {
	testcases := []struct {
		name     string
		expected bool
	}{
		{"test0", true},
	}

	bs := ByteStream{}

	for _, e := range testcases {
		got := bs.IsEmpty()
		if got != e.expected {
			t.Errorf("Test case %s failed. Expected %v, but got %v\n", e.name, e.expected, got)
		}
	}
}
