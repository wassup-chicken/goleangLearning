package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	first string
	ltk   bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		first: "ETHAN HAWK",
		ltk:   true,
	}

	p2 := person{
		first: "Jen",
		last:  "Money",
		age:   27,
	}

	fmt.Printf("%T \t %#v\n", sa1, sa1)
	fmt.Println(p2)

}
