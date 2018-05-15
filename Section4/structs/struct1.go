package main

import "fmt"

type Foo struct {
	A int
	b string
}

func main() {
	f := Foo{}
	fmt.Println(f)

	g := Foo{10, "Hello"}
	fmt.Println(g)

	h := Foo{
		b: "Goodbye",
	}
	fmt.Println(h)

	h.A = 1000
	fmt.Println(h.A)
}
