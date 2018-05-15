package main

import "fmt"

type Foo struct {
	A int
	B string
}

func (f Foo) String() string {
	return fmt.Sprintf("A: %d and B: %s", f.A, f.B)
}

func (f *Foo) Double() {
	f.A = f.A * 2
}

func main() {
	f := Foo{
		A: 10,
		B: "Hello",
	}
	fmt.Println(f.String())
	f.Double()
	fmt.Println(f.String())
}
