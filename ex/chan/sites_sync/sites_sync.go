// Get content type of sites
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	dt := time.Now()
	t := dt.Format("15:04:05.999999999")

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s ::: %s -> %s\n", t, url, ctype)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
		"https://en.wikipedia.org/",
		"https://courses.calhoun.io/",
		"https://gobyexample.com/",
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			returnType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}
