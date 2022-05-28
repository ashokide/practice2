package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
	City string
}

func main() {
	user := &User{Name: "Ashok", Age: 20, City: "Erode"}
	data, _ := json.Marshal(user)
	fmt.Println(string(data))
}
