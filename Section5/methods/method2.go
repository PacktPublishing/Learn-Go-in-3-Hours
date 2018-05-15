package main

import "fmt"

type Foo struct {
	A int
	B string
}

func (f Foo) fieldCount() int {
	return 2
}

func (f Foo) String() string {
	return fmt.Sprintf("%d fields: A: %d and B: %s", f.fieldCount(), f.A, f.B)
}

func (f *Foo) Double() {
	f.A = f.A * 2
}

type Bar struct {
	C bool
	Foo
}

func (b Bar) String() string {
	return fmt.Sprintf("%s and C: %v", b.Foo.String(), b.C)
}

func (b Bar) fieldCount() int {
	return 3
}

func main() {
	f := Foo{
		A: 10,
		B: "Hello",
	}
	b := Bar{
		C:   true,
		Foo: f,
	}
	fmt.Println(b.String())
	b.Double()
	fmt.Println(b.String())
}
