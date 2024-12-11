package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

func main() {
	fmt.Println("Hello!")
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)
}
