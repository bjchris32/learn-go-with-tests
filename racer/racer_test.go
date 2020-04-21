package racer

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.co.uk"

	expect := fastURL
	got := Racer(slowURL, fastURL)

	if got != expect {
		t.Errorf("got %q, expect %q", got, expect)
	}
}
