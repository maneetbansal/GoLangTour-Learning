package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Cache to keep track of fetched URLs
var cache map[string]bool = make(map[string]bool)

// Mutex for synchronized access to cache
var cache_mu sync.Mutex

// Using WG for parallelization, you can use channel as well
var wg sync.WaitGroup

func hasUrlFetched(url string) bool {
	cache_mu.Lock()
	defer cache_mu.Unlock()
	ret := false
	if cache[url] {
		ret = true
	}
	return ret
}

func cacheUrl(url string) {
	cache_mu.Lock()
	defer cache_mu.Unlock()
	cache[url] = true
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// Will release from WaitGroup at end of function call
	defer wg.Done()
	if depth <= 0 {
		return
	}
	// If Url has already been fetched, just return
	if hasUrlFetched(url) {
		return
	}

	// Cache the URL before facthing
	cacheUrl(url)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		// debug statement
		//fmt.Println("Crawling url: ", u)
		// Adding to WaitGroup
		wg.Add(1)
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher)
	wg.Wait()
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
