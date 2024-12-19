package main

import (
	"fmt"
)

func main() {
	//array
	a := [3]int{42, 43, 44}
	fmt.Println(a)

	b := [...]string{"hello", "gophers"}
	fmt.Println(b)

	var c [3]int
	c[0] = 7
	c[1] = 8
	fmt.Println(c)

	{
		var ni [7]int

		ni[0] = 42
		fmt.Printf("%#v \t\t\t\t %T\n", ni, ni)
	}
}
