package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	restChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			restChannel <- result{u, wc(u)}
		}(url)
	}

	for i:= 0;i<len(urls); i++{
		r:= <-restChannel
		results[r.string] = r.bool
	}

	return results
}
