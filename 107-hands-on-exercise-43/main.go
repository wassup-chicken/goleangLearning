package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range x {
		fmt.Printf("%v \t %T \t %v\n", v, v, i)
	}

	fmt.Printf("%T \t %#v\n", x, x)
	fmt.Printf("%T \t %v\n", x, x)

}

/*
Using a COMPOSITE LITERAL:
create a SLICE OF TYPE int
assign 10 values
42, 43, 44, 45, 46, 47, 48, 49, 50, 51
Range over the slice and print the values.
printout the value and the type
*/
