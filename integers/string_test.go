package integers

import (
	"strings"
	"testing"
)

func TestStringCount(t *testing.T) {
	got := strings.Count("wtffffff", "ff")
	want := 3
	if got != want {
		t.Errorf("got %q want %q", got, want)

	}
}
func TestFuck(t *testing.T) {
	got := strings.Trim("fwtfff", "tf")
	want := "w"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
