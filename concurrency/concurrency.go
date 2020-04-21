package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		// we need the argument u, so that the goroutine would not share the same url variable
		go func(u string) {
			// whenever it finishes, it puts the result into the channel
			resultChannel <- result{u , wc(u)}
		}(url)
    }

	for i := 0; i < len(urls); i++ {
		// read the result from channel
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
