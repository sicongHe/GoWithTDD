package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsitechecker(_ string) bool {
	return true
}

func slowWebsitechecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsite(t *testing.T) {
	urls := []string{
		"www.1.com",
		"www.2.com",
		"www.3.com",
	}
	got := CheckWebsite(mockWebsitechecker, urls)
	want := map[string]bool{
		"www.1.com": true,
		"www.2.com": true,
		"www.3.com": true,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v,want %#v", got, want)
	}
}

func BenchmarkName(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "url"
	}
	for i := 0; i < b.N; i++ {
		CheckWebsite(slowWebsitechecker, urls)
	}
}
