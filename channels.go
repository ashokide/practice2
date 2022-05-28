package main

import (
	"fmt"
)

func main() {
	c := make(chan string,10)
	fmt.Println("Size of Channel : ", cap(c))
}
