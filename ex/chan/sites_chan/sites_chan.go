// Get content type of sites
package main

import (
	"fmt"
	"net/http"
	"time"
)

func returnType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	dt := time.Now()
	t := dt.Format("15:04:05.999999999")

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	res := fmt.Sprintf("%s ::: %s -> %s", t, url, ctype)
	return res, nil
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

	ch := make(chan string)

	for _, url := range urls {
		go func(url string) {
			result, err := returnType(url)
			if err != nil {
				fmt.Printf("error: %s\n", err)
			}
			ch <- result
		}(url)
	}

	for range urls {
		val := <-ch
		fmt.Println(val)
	}
}
