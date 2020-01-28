package main

import (
	"fmt"
	"sync"
)

type Crawler struct {
	crawled map[string]bool
	mux     sync.Mutex
}

func New() *Crawler {
	return &Crawler{crawled: make(map[string]bool)}
}

func (crawler *Crawler) Crawl(url string, depth int, fetcher fakeFetcher) {
	var wg sync.WaitGroup

	if depth <= 0 || crawler.visit(url) {
		return
	}

	body, urls, e := fetcher.Fetch(url)
	if e != nil {
		fmt.Println(e)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			crawler.Crawl(u, depth-1, fetcher)
		}(u)
	}
	wg.Wait()
	return
}

func (crawler *Crawler) visit(url string) bool {
	crawler.mux.Lock()
	defer crawler.mux.Unlock()

	_, ok := crawler.crawled[url]
	if ok {

		return true
	}
	crawler.crawled[url] = true
	return false
}

func main() {
	crawler := New()
	crawler.Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
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

// fetcher is a populated fakeFetcher.
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
