package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string { "test": "this is just a test" }

	got := Search(dictionary, "test")
	expected := "this is just a test"

	if got != expected {
		t.Errorf("got %q expected %q given, %q", got, expected, "test")
	}
}
