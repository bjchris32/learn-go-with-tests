package main

import "testing"
import "bytes"
import "reflect"


func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountdownOperationsSpy{})

		got := buffer.String()
		expected := `3
2
1
Go!`

		if got != expected {
			t.Errorf("got %q expected %q", got, expected)
		}
	})

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
