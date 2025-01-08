package main

import "fmt"

func main() {
	defer bar()
	foo()
	bar()
}

//func (r receiver) identifier(p parameter(s)) (return(s)) { code }

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
