package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("sum 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d and given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	n1 := []int{1, 2, 3}
	n2 := []int{1, 2, 3, 4, 5}

	got := sumAll(n1, n2)
	want := []int{6, 15}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("with normal slices", func(t *testing.T) {
		got := sumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("with empty slices", func(t *testing.T) {
		got := sumAllTails([]int{})
		want := []int{0}
		checkSums(t, got, want)
	})

}
