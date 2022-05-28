package main

import "fmt"

func main() {
	distinct := make(map[rune]int)
	var str string = "Asshok"
	// fmt.Scanln(&str)

	for _, c := range str {
		// if _, pre := distinct[c]; pre {
		// 	fmt.Println("Already Present - ", distinct[c])
		// } else {
		distinct[c]++
		// }
	}
	for k, c := range distinct {
		fmt.Println(string(k), " - ", c)
	}
}
