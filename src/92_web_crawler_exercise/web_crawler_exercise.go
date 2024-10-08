package main

import (
	"fmt"
	"time"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

var cachedUrls = make(map[string]bool)
var mutex sync.Mutex

func Crawl(url string, depth int, fetcher Fetcher) {
	mutex.Lock()
	_, cached := cachedUrls[url]
	mutex.Unlock()
	if cached {
		return
	} else {
		mutex.Lock()
		cachedUrls[url] = true
		mutex.Unlock()
	}

	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)
	}
}

func main() {
	const maxCrawlTime = time.Millisecond*500
	Crawl("https://golang.org/", 4, fetcher)
	time.Sleep(maxCrawlTime)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}

	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}