package main

import "fmt"

func reallyNil() error {
	var e error
	fmt.Println("e is nil:", e == nil)
	return e
}

type MyError struct {
	A       int
	B       int
	Message string
}

func (me *MyError) Error() string {
	return fmt.Sprintf("values %d and %d produced error %s", me.A, me.B, me.Message)
}

func notReallyNil() error {
	var me *MyError
	fmt.Println("me is nil:", me == nil)
	return me
}

func main() {
	e := reallyNil()
	me := notReallyNil()
	fmt.Println("in main, e is nil:", e == nil)
	fmt.Println("in main, me is nil:", me == nil)
}
