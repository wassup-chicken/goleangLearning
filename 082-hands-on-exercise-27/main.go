package main

import "fmt"

//break using for loop without any condition
func main() {
	x := 20
	for {
		fmt.Printf("%v \n", x)
		if x < 0 {
			break
		}
		x--
	}
}
