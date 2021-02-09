package util

import "testing"

func TestBytesToString(t *testing.T) {
	var b []byte = []byte("abc")
	s := BytesToString(b)
	if s != string(b) {
		t.Fatal("BytesToString error")
	}
	b[1] = 'c'
	if s != "acc" {
		t.Fatal("BytesToString error")
	}
}
