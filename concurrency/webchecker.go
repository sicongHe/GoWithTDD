package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
	ret := make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		go func(u string) { resultChannel <- result{u, wc(u)} }(url)
	}
	for i := 0; i < len(urls); i++ {
		resultFromChannel := <-resultChannel
		ret[resultFromChannel.string] = resultFromChannel.bool
	}
	return ret
}
