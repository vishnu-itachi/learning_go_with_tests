package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"

		assertMessage(t, got, want)
	})

	t.Run("saying hello when empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("jay", "Spanish")
		want := "Hola, jay"

		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
