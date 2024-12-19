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

	for i, v := range xs {
		fmt.Printf("%v - value %v\n", i, v)
	}

}
