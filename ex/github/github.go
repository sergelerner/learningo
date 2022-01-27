package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name        string `json:"name"`
	PublicRepos int    `json:"public_repos"`
}

func userInfo(login string) (*User, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", login)

	// HTTP call
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode JSON
	user := &User{}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

func main() {
	user, err := userInfo("sergelerner")
	if err != nil {
		log.Fatalf("Fail fetching user from github")
	}
	fmt.Printf("user %v#\n", user)
}
