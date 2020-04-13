package main

import "testing"
import "bytes"

func TestGreet(t * testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	expected := "Hello, Chris"

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}
