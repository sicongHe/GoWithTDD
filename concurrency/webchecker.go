package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
	ret := make(map[string]bool)
	resultChannl := make(chan result)
	for _, url := range urls {
		go func(u string) { resultChannl <- result{url, wc(url)} }(url)
	}
	for i := 0; i < len(urls); i++ {
		resultFromChannl := <-resultChannl
		ret[resultFromChannl.string] = resultFromChannl.bool
	}
	return ret
}
