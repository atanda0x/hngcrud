package main

import "testing"

func Test(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("atanda0x")
		want := "Hello, atanda0x"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello  World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
