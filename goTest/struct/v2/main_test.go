package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}

	got := Perimeter(rectangle)
	want := 40.0

	assertMessage(t, got, want)
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{12.0, 6.0}

	got := Area(rectangle)
	want := 72.0

	assertMessage(t, got, want)

}

func assertMessage(t testing.TB, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("got %.2f but want %.2f", got, want)
	}
}
