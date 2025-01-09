package main

import (
	"fmt"
)

func intDelta(n *int) {
	*n = 43
}

func sliceDelta(ii []int) {
	ii[0] = 99
}

func mapDelta(md map[string]int, s string) {
	md[s] = 33
}

func main() {
	a := 42
	fmt.Println(a)

	intDelta(&a)
	fmt.Println(a)

	//slice
	xi := []int{1, 2, 3, 4}
	fmt.Println(xi)
	sliceDelta(xi)
	fmt.Println(xi)

	//map
	m := make(map[string]int)
	m["James"] = 32
	fmt.Println(m["James"])
	mapDelta(m, "James")
	fmt.Println(m["James"])

}

// if value that's being passed is a reference type, it's like a pointer to an int
// it's referencing a memory location or it's a lice, memory, etc..
//Pointers, Slices, Maps, Channels, Functions, Interfaces
