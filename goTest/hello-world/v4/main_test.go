package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying Hello", func(t *testing.T) {
		got := Hello("0x")
		want := "Hello, 0x"

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
}
