package main

import "fmt"

func main() {
	foo()

	//anonymous
	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(s string) { fmt.Println("YOOOO", s) }("Todd")
}

func foo() {
	fmt.Println("Foo ran")
}

//func (r receiver) identifier(p parameter(s)) (return(s)) { code }

// an anonymous function
// func(p parameter(s)) (r return(s)) { code }
