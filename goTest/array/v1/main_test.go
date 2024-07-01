package main

import "testing"

func TestArray(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	assertMessage(t, got, want, numbers[:])
}

func assertMessage(t *testing.T, got, want int, numbers []int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d but want %d and was given %v", got, want, numbers)
	}
}
