package main

import (
	"fmt"
)

func main() {
	//slice append

	xi := []int{42, 43, 44}
	fmt.Println(xi)
	fmt.Println("---------------")
	// variadic parameter
	xi = append(xi, 33, 45, 46, 47, 99, 77)
	fmt.Println(xi)
	fmt.Println("---------------")

	for i := 0; i < 10; i++ {
		xi = append(xi, i)
	}
	fmt.Println(xi)

}
