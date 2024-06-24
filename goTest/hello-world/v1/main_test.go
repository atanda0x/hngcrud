package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello-World"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
