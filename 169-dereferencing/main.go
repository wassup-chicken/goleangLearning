package main

import (
	"fmt"
)

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)

	//ampersand gives the address
	y := &x
	fmt.Printf("%v\t%T\n", y, y)
	fmt.Println(*y)

	//asterisk dereferences the address and gives the value at that address
	*y = 43
	fmt.Println(x)
	fmt.Println(y)

}
