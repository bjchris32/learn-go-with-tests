package racer

import(
	"testing"
	"net/http/httptest"
	"time"
	"net/http"
)

func TestRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	expect := fastURL
	got := Racer(slowURL, fastURL)

	if got != expect {
		t.Errorf("got %q, expect %q", got, expect)
	}

	slowServer.Close()
	fastServer.Close()
}
