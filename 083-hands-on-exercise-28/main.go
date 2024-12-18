package main

import "fmt"

//use modulus and continue in a loop to print all ODD numbers
func main() {
	x := 0
	for x < 100 {
		// print i when it is an odd number
		if x%2 != 0 {
			fmt.Println(x)

		}
		x++
	}
}
