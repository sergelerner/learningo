package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type Client struct {
	send chan Message
}

func (client *Client) write() {
	for msg := range client.send {
		// TODO socket.sendJSON(msg)
		fmt.Printf("%#v\n", msg)
	}
}

func (client *Client) subscribeChannles() {
	for {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		client.send <- Message{"channel add", ""}
	}
}

func (client *Client) subscribeMessages() {
	for {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		client.send <- Message{"message add", ""}
	}
}

func NewClient() *Client {
	return &Client{
		send: make(chan Message),
	}
}

func main() {
	client := NewClient()
	go client.subscribeChannles()
	go client.subscribeMessages()
	client.write()
}
