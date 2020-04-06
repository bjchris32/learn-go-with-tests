package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary { "test": "this is just a test" }

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		expected := "this is just a test"

		assertStrings(t, got, expected)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		expected := "could ont find the word"

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, err.Error(), expected)
	})
}

func assertStrings(t *testing.T, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}
