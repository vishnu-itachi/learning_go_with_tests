package arrays

import "testing"


import "reflect"

func TestSum(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)

	want := 15

	if got != want {
		t.Errorf("Expected %d but got %d", want, got)
	}

}

func TestSumSlice(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6}

	got := SumSlice(numbers)
	want := 21

	if got != want {
		t.Errorf("Expected %d but got %d", want, got)
	}
}


func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}