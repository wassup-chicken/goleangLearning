package main

import (
	"fmt"

	"math/rand"
)

// use modulus and continue in a loop to print all ODD numbers
func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	age := m["James"]
	fmt.Println("The age of Bond", age)
	if v, ok := m["James"]; ok {
		fmt.Println("There is Bond, and here is the Bond's age ", v)
	}

	age = m["Q"]
	fmt.Println("The age of Q", age)

	//comma okay
	if v, ok := m["Q"]; !ok {
		fmt.Println("There is no Q, and here is the zero value of an int ", v)
	}

	fmt.Println("---------------------------")

	c := 1
	for i := 0; i < 100; i++ {
		//statement statement idiom
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("iteration %v \t  total count %v \t x is %v\n", i, c, x)
			c++
		}
	}
}
