package racer

import (
	"time"
	"fmt"
	"net/http"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	// wait multiple channel at the same time
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	// this case prevents the program from blocking forever
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	// this channel is used as signal to notify that the work is done
	// it does not matter what is in the channel.
	// Using struct as the smallest memory allocation.
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		// notify the system that http.Get is done
		close(ch)
	}()

	return ch
}
