package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	//this is anonymous struct, used when using just once
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "abc",
		last:  "def",
		age:   27,
	}

	p2 := person{
		first: "Jen",
		last:  "Money",
		age:   27,
	}

	fmt.Printf("%T \n", p1)
	fmt.Printf("%T \n", p2)

}
