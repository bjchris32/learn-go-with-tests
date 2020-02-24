package main

import "testing"

// *testing.T is the hook into the testing framework
func TestHello(t *testing.T) {
	assertMessage := func(t *testing.T, got string, expected string) {
		t.Helper() // for tracing the failure line number in the caller of this assertion
		if got != expected {
			t.Errorf("got %q expected %q", got, expected)
		}
	}

	t.Run("greeting to people", func(t *testing.T) {
		got := Hello("Chris", "")
		expected := "Hello, Chris"
		assertMessage(t, got, expected)
	})


	t.Run("greeting the world when there is no argument", func(t *testing.T) {
		got := Hello("", "")
		expected := "Hello, World"
		assertMessage(t, got, expected)
	})

	t.Run("greeting in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		expected := "Hola, Elodie"
		assertMessage(t, got, expected)
	})
}
