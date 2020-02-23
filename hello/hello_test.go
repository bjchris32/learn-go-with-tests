package main

import "testing"

// *testing.T is the hook into the testing framework
func TestHello(t *testing.T) {
	got := Hello("Chris")
	expected := "Hello, Chris"

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}
