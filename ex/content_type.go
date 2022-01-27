package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	ctype := resp.Header.Get("Content-Type")
	if ctype == "" {
		return "", fmt.Errorf("Cant find content type")

	}
	return ctype, nil
}

func ct() {
	thetype, err := contentType("https://gobyexample.co/")
	if err != nil {
		fmt.Printf("Failed getting content type %s\n", err)
	}
	fmt.Println(thetype)
}
