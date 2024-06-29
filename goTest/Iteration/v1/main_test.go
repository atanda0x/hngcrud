package main

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("0x")
	expected := "0x0x0x0x0x"

	assertEror(t, expected, repeated)
}

func assertEror(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("expected %q but got %q", got, want)
	}
}
