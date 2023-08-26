package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("sayin hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		assertCorrectMessage(got, want, t)
	})

	t.Run("say 'Hello World' when sting is empty", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		assertCorrectMessage(got, want, t)
	})

}

func assertCorrectMessage(got, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
