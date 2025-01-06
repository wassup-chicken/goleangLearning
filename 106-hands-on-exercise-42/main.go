package main

import "fmt"

func main() {
	// create an array
	x := [5]int{}
	//assign values to each index position
	for i := 0; i < 5; i++ {
		x[i] = i
	}

	for i, v := range x {
		fmt.Printf("%v - %T - %v\n", v, v, i)
	}

}

/*
Using a COMPOSITE LITERAL:
create an Array which holds 5 values of type int
assign values of each index position.
Range over the array and print the values.
printout the value and the type
*/
