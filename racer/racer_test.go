package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("快慢请求竞争", func(t *testing.T) {
		slowServer := getServer(20 * time.Millisecond)
		fastServer := getServer(0)
		defer slowServer.Close()
		defer fastServer.Close()
		slowUrl := slowServer.URL
		fastUrl := fastServer.URL
		want := fastUrl
		got, err := Race(fastUrl, slowUrl, 10*time.Second)
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}
		if got != want {
			t.Errorf("got %s,want %s", got, want)
		}
		slowServer.Close()
		fastServer.Close()
	})
	t.Run("超时错误", func(t *testing.T) {
		A := getServer(20 * time.Millisecond)
		defer A.Close()

		_, err := Race(A.URL, A.URL, 1*time.Millisecond)
		assertError(t, err, "请求超时")
	})
}

func getServer(sleep time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(sleep)
		w.WriteHeader(http.StatusOK)
	}))
}

func assertError(t *testing.T, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("need an err but got nil")
	}
	if got.Error() != want {
		t.Errorf("Got %s, Want %s", got.Error(), want)
	}
}
