package integers

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 3)
	expected := "aaa"
	if repeated != expected {
		t.Errorf("expected %q repeated %q", expected, repeated)

	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	out := Repeat("a", 3)
	fmt.Println(out)
	// Output: aaa
}
