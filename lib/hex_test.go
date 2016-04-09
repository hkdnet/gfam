package lib

import (
	"bytes"
	"testing"
)

func TestHexToBytes(t *testing.T) {
	if got, want := hexToBytes("a"), []byte{10}; !bytes.Equal(got, want) {
		t.Errorf("want: %v\ngot: %v", want, got)
	}
	if got, want := hexToBytes("0xa"), []byte{10}; !bytes.Equal(got, want) {
		t.Errorf("want: %v\ngot: %v", want, got)
	}
	if got, want := hexToBytes("aaa"), []byte{10, 170}; !bytes.Equal(got, want) {
		t.Errorf("want: %v\ngot: %v", want, got)
	}
}
