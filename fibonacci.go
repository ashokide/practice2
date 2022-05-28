package main

import (
	"fmt"
)

func Fibonacci(n int64) int64 {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	var n int64 = 50
	result := Fibonacci(n)
	fmt.Println("Fibonacci of ", n, " is : ", result)
}
