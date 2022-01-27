package main

import (
	"fmt"
	"log"
)

type Point struct {
	x int
	y int
}

type Square struct {
	center Point
	length int
}

func (s *Square) Move(dx int, dy int) {
	s.center.x = dx
	s.center.y = dy
}

func (s *Square) Area() int {
	area := s.length * s.length
	return area
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length < 0 {
		return nil, fmt.Errorf("length cant be less then 0")
	}
	p := Point{x, y}
	s := &Square{center: p, length: length}
	return s, nil
}

func main() {
	square, err := NewSquare(5, 5, 5)
	if err != nil {
		log.Fatalf("Failed creating a square")
	}
	square.Move(10, 10)
	fmt.Printf("position of our square %+v\n", square.center)
	fmt.Println("the are of our square", square.Area())
}
