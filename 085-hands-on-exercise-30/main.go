package main

import "fmt"

//use modulus and continue in a loop to print all ODD numbers
func main() {
	xi := []int{42, 43, 44, 45, 46, 47}

	for i, v := range xi {
		fmt.Printf("Index %v \t value %v\n", i, v)
	}
}
