package concurrency

import "time"

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		// we need the argument u, so that the goroutine would not share the same url variable
		go func(u string) {
			results[u] = wc(u)
		}(url)
    }

    // wait for the goroutines done
    time.Sleep(2 * time.Second)

    return results
}
