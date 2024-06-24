package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Atanda0x")
	want := "Hello, Atanda0x"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
