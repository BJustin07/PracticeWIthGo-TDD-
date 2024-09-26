package arraysAndSlices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Using array here, fixed number of elements", func(t *testing.T) {
		numbersArray := []int{1, 2, 3, 4, 5}
		sum := Sum(numbersArray)
		if sum != 15 {
			t.Errorf("Sum was incorrect, got: %d, want: %d., given: %v", sum, 15, numbersArray)
		}
	})
	t.Run("Using slice here, any number of elements", func(t *testing.T) {
		numbersArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		sum := Sum(numbersArray)
		if sum != 55 {
			t.Errorf("Sum was incorrect, got: %d, want: %d., given: %v", sum, 15, numbersArray)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{4, 5, 6})
	want := []int{6, 15}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SumAll was incorrect, got: %v, want: %v.", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("SumAllTails was incorrect, got: %v, want: %v.", got, want)
		}
	}
	t.Run("Sum of the tails of X number of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{5, 11}
		checkSums(t, got, want)
	})
	t.Run("Sum of the tails of one empty slice among slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}
		checkSums(t, got, want)
	})
}
