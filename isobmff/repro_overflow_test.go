package isobmff

import "testing"

// TestParseContainerHugePayloadLen reproduces the "makeslice: len out of range"
// panic at container.go:49: a largesize box whose Size is near 2^64 makes
// PayloadLen() return a value > MaxInt64, which overflows the int64 cast in
// ParseContainer and produces a negative make length.
func TestParseContainerHugePayloadLen(t *testing.T) {
	// largesize marker (size == 1), type "ftyp", then a 64-bit size near 2^64.
	data := []byte{
		0x00, 0x00, 0x00, 0x01, // size == 1 -> largesize form
		'f', 't', 'y', 'p',
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // Size = 2^64-1
	}
	_, err := ParseContainer(data)
	if err == nil {
		t.Fatal("ParseContainer: expected error for oversized payload, got nil")
	}
}
