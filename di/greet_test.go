package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "test")
	got := buffer.String()
	want := "hello test"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
