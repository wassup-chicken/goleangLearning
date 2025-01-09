package main

import "fmt"

func main() {
	foo()

	x := func() {
		fmt.Println("Anonymous func ran")
	}

	y := func(s string) { fmt.Println("YOOOO", s) }

	x()
	y("Todd")
}

func foo() {
	fmt.Println("Foo ran")
}

//func (r receiver) identifier(p parameter(s)) (return(s)) { code }
