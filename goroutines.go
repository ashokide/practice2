package main

import (
	"fmt"
	"time"
)

type Book struct {
	authorId   string
	authorName string
}

func printDetails(b Book) {
	fmt.Println("Author ID - ", b.authorId)
	fmt.Println("Author Name - ", b.authorName)
}

func main() {
	obj := Book{
		"19ITR013",
		"Ashok N",
	}

	obj1 := Book{
		"19ITR014",
		"Alvin",
	}

	go printDetails(obj)
	time.Sleep(time.Second * 1)
	go printDetails(obj1)
	time.Sleep(time.Second * 1)
}
