package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	bob := `{ "name": "Bob", "age": 30}`
	var b Person
	json.Unmarshal([]byte(bob), &b)
	fmt.Println(b)
	bob2, _ := json.Marshal(b)
	fmt.Println(string(bob2))
}
