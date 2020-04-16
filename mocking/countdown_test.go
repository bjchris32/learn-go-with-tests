package main

import "testing"
import "bytes"
import "reflect"


func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	expected := `3
2
1
Go!`

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, expected 4 got %d", spySleeper.Calls)
	}

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy {}
		Countdown(spySleepPrinter, spySleepPrinter)

		expected := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(expected, spySleepPrinter.Calls) {
			t.Errorf("expected calls %v got %v", expected, spySleepPrinter.Calls)
		}
	})
}
