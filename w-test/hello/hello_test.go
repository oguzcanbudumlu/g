package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to kamil", func(t *testing.T) {
		got := Hello("kamil", "")
		want := "hello kamil"

		assertCorrectMessage(t, got, want)
	})

	t.Run("'hello world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "hello World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Kamil", "Spanish")
		want := "Hola Kamil"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
