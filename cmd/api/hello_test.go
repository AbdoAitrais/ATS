package main

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Run("Saying Hello to people", func(t *testing.T) {
		got := Hello("Sami")
		want := "Hello Sami"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Default salute for empty name", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	//t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
