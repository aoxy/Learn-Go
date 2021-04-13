package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher, r *SafeRecorder) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
	// 下面并没有实现上面两种情况：
	if depth <= 0 {
		return
	}
	r.Tag(url)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if !r.Exist(u) {
			go Crawl(u, depth-1, fetcher, r)
		}
	}
	return
}

func main() {
	r := &SafeRecorder{visited: make(map[string]bool)}
	go Crawl("https://golang.org/", 4, fetcher, r)

	var str string
	fmt.Scan(&str)
	// 这两句代码是等待输入的意思，在这里用来阻止主线程关闭的。如果没有这两句的话，会发现我们的程序瞬间就结束了，而且什么都没有输出。这是因为主线程关闭之后，所有开启的goroutine都会强制关闭，他还没有来得及输出，就结束了。
	// 可以使用channel优化下
}

// fakeFetcher 是返回若干结果的 Fetcher。
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

type SafeRecorder struct {
	visited map[string]bool
	mux     sync.Mutex
}

func (r *SafeRecorder) Tag(key string) {
	r.mux.Lock()
	r.visited[key] = true
	r.mux.Unlock()
}

func (r *SafeRecorder) Exist(key string) bool {
	r.mux.Lock()
	defer r.mux.Unlock()
	return r.visited[key]
}

// fetcher 是填充后的 fakeFetcher。
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
