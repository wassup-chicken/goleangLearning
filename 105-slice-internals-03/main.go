package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := a

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

	x := 10
	y := x
	x = 20
	fmt.Println("x ", x)
	fmt.Println("y ", y)
	fmt.Println("-----------------")
}
