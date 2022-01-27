package main

import (
	"fmt"
	"strings"
)

func words() {
	count := map[string]int{}
	text := `
		Needls and pins
		Needls and pins
		Sew me a sail
		To catch me the wind
	`

	list := strings.Fields(strings.ToLower(text))

	for _, val := range list {
		current := count[val]
		count[val] = current + 1
	}

	fmt.Println(count)
}
