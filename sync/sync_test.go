package main

import "testing"

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

}

func assertCounter(t *testing.T, got Counter, expected int) {
	t.Helper()
	if got.Value() != expected {
		t.Errorf("got %d, expected %d", got.Value(), expected)
	}
}
