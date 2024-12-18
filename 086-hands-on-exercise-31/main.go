package main

import "fmt"

//use modulus and continue in a loop to print all ODD numbers
func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Printf("Name %v \t Age %v\n", k, v)
	}
}
