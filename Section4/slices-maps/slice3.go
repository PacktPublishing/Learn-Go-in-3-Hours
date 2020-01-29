package main

import "fmt"

func main() {
	uniHello := "👋 🌍"
	bytes := []byte(uniHello)
	fmt.Println(bytes)
	runes := []rune(uniHello)
	fmt.Println(runes)
	runes[1] = 'a'
	fmt.Println(runes)
	fmt.Println(uniHello)
}
