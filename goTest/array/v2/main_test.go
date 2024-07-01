package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		got := Sum(numbers)
		want := 45

		assertMessage(t, got, want, numbers[:])
	})

	t.Run("Collection of any size(slice)", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		got := Sum(numbers)
		want := 10

		assertMessage(t, got, want, numbers)
	})
}

func assertMessage(t *testing.T, got, want int, numbers []int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d but want %d was given %v", got, want, numbers)
	}
}
