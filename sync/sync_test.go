package main

import "testing"
import "sync"


func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		expectedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(expectedCount)

		for i := 0; i < expectedCount; i++ {
			// the go routine should dereference the wait group address to get the value
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg) // the go routine needs to take wait group's address as arguemnt
		}

		wg.Wait()

		assertCounter(t, counter, expectedCount)
	})
}

func assertCounter(t *testing.T, got Counter, expected int) {
	t.Helper()
	if got.Value() != expected {
		t.Errorf("got %d, expected %d", got.Value(), expected)
	}
}
