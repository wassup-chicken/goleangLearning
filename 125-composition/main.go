package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type foo int

func main() {
	var a foo = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Bond",
		last:  "james",
		age:   33,
	}

	fmt.Println(p1)

	fmt.Printf("%T \t %#v\n", p1, p1)

	p2 := person{
		first: "Yo",
		last:  "Hi",
		age:   34,
	}

	fmt.Println(p2)
	fmt.Printf("%T \t %#v\n", p2, p2)

}
