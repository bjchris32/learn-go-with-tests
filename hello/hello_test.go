package main

import "testing"

// *testing.T is the hook into the testing framework
func TestHello(t *testing.T) {
	t.Run("greeting to people", func(t *testing.T) {
		got := Hello("Chris")
		expected := "Hello, Chris"

		if got != expected {
			t.Errorf("got %q expected %q", got, expected)
		}
	})


	t.Run("greeting the world when there is no argument", func(t *testing.T) {
		got := Hello("")
		expected := "Hello, World"

		if got != expected {
			t.Errorf("got %q expected %q", got, expected)
		}
	})
}
