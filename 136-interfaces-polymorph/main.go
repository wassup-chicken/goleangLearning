package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I am", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I a secret agent", sa.first)
}

// inteface: hey if you got these methods, then you're my type.
// person has speak method and SA have this method.
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}

	p2 := person{
		first: "Yooo",
	}

	saySomething(sa1)
	saySomething(p2)
}

//func (r receiver) identifier(p parameter(s)) (return(s)) { code }
