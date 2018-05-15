package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "👋 🌍"
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
}
