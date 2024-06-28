package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("say Hello to people", func(t *testing.T) {
		got := Hello("0x", "English")
		want := "Hello, 0x"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say Hello World when the string is empty", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say Hello in Spanish", func(t *testing.T) {
		got := Hello("0x", "Spanish")
		want := "Hola, 0x"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say Hello in French", func(t *testing.T) {
		got := Hello("0x", "French")
		want := "Bonjour, 0x"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say Hello in Yoruba", func(t *testing.T) {
		got := Hello("0x", "Yoruba")
		want := "Bawoni, 0x"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
