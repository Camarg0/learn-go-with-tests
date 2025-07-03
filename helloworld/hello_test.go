package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Matheus", "")
		want := "Hello, Matheus"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Matheus", "Spanish")
		want := "Hola, Matheus"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Matheus", "French")
		want := "Bonjour, Matheus"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("Matheus", "Portuguese")
		want := "Ola, Matheus"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // numero da linha do erro reporta na chamada da função, não aqui dentro
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
