package main

import (
	"fmt"
)

func main() {
	//array
	xs := []string{"a", "b", "c", "d", "ef", "g", "h", "j"}
	fmt.Println(xs)
	fmt.Println(len(xs))
	fmt.Printf("%T\n", xs)

	for _, v := range xs {
		fmt.Printf("%v\n", v)
	}

	fmt.Println("------------")
	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])

	fmt.Println("------------")
	for i := 0; i < len(xs); i++ {
		fmt.Println(xs[i])
	}
}
