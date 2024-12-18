package main

import "fmt"

//use modulus and continue in a loop to print all ODD numbers
func main() {
	x := 0
	for x < 5 {
		// print i when it is an odd number
		y := 0
		fmt.Println("this is the outer")
		for y < 5 {
			fmt.Printf("Out at %v \t and inner at %v\n", x, y)
			y++
		}
		x++
	}
}
