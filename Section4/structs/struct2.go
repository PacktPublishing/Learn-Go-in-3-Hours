package main

import "fmt"

type Foo struct {
	A int
	b string
}

func main() {
	f := Foo{
		A: 20,
	}
	var f2 Foo
	f2 = f
	f2.A = 100
	fmt.Println(f2.A)
	fmt.Println(f.A)

	var f3 *Foo = &f
	f3.A = 200
	fmt.Println(f3.A)
	fmt.Println(f.A)
}
