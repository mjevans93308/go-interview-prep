package main

import (
	"fmt"

	"github.com/mjevans93308/go-interview-prep/datastructures"
)

func main() {
	var s datastructures.Stack

	s.Push(1)
	s.Push(3)
	s.Push(2)

	fmt.Println(s)
}
