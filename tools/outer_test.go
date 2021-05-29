package tools

import "testing"

func TestHi(t *testing.T) {
	got := Hi("Sicong")
	want := "hi Sicong!\n"
	if got != want {
		t.Errorf("got %q", got)
	} else {
		t.Logf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	got := Hello()
	want := "hello!\n"
	if got != want {
		t.Errorf("got %q", got)
	}
}
