package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello with name", func(t *testing.T) {
		got := Hello("bishes", "English")
		want := "Hello, bishes"
		assertMessage(t, got, want)
	})
	t.Run("Saying hello world if name has no value", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertMessage(t, got, want)
	})
	t.Run("In tagalog", func(t *testing.T) {
		got := Hello("Enteng", "Tagalog")
		want := "Kamusta, Enteng"
		assertMessage(t, got, want)
	})
	t.Run("In bisaya", func(t *testing.T) {
		got := Hello("Enteng", "Bisaya")
		want := "Maupay, Enteng"
		assertMessage(t, got, want)
	})
}

func assertMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
