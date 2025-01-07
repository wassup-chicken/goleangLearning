package main

import "fmt"

func main() {
	foo()

	bar("James")

	s := aloha("Bond")
	fmt.Println(s)

	s1, n := dogYears("Woof", 40)
	fmt.Println(s1, n)
}

// no params no returns
func foo() {
	fmt.Println("I am from foo")
}

// 1 param, no returns
func bar(s string) {
	fmt.Println("My name is", s)
}

// 1 param, 1 return
func aloha(s string) string {
	return fmt.Sprint("My name is again ", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, "is this old in dog years"), age
}

//func (r receiver) identifier(p parameter(s)) (return(s)) { code }
