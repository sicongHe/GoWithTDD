package main

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := sum(numbers)
	want := 15
	if got != want {
		t.Errorf("got %d want %d and given %v", got, want, numbers)
	}

}
