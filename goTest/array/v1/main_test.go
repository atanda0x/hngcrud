package main

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(2, 3)
	expected := 5

	assertMessage(t, sum, expected)
}

func assertMessage(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}
