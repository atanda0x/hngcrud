package main

import "testing"

func Test(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("atanda0x")
		want := "Hello, atanda0x"

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Say 'Hello  World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
}
