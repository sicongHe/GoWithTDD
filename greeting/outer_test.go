package greeting

import "testing"

func TestHi(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hi to people", func(t *testing.T) {
		got := Hi("Sicong", "")
		want := "hi Sicong"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say hi to nobody", func(t *testing.T) {
		got := Hi("", "")
		want := "hi gopher"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say hi in Chinese", func(t *testing.T) {
		got := Hi("思聪", "Chinese")
		want := "你好 思聪"
		assertCorrectMessage(t, got, want)
	})

}

func TestHello(t *testing.T) {
	got := Hello()
	want := "hello "
	if got != want {
		t.Errorf("got %q", got)
	}
}
