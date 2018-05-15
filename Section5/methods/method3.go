package main

import "fmt"

type myInt int

func (mi myInt) isEven() bool {
	return mi%2 == 0
}

func (mi *myInt) Double() {
	*mi = *mi * 2
}

func main() {
	m := myInt(10)
	fmt.Println(m)
	fmt.Println(m.isEven())
	m.Double()
	fmt.Println(m)
}
