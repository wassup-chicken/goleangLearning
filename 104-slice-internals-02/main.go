package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := make([]int, 6)
	// b := []int{0, 0, 0, 0, 0, 0}
	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("-----------------")

	// copy (destination, source)
	copy(b, a)

	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("-----------------")

	a[0] = 7
	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("-----------------")

	b[1] = 999
	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("-----------------")
}
