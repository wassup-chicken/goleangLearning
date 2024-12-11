package main

import (
	"fmt"

	"github.com/wassup-chicken/puppy"
)

func main() {
	fmt.Println(puppy.Bark())

	s1 := puppy.Bark()

	fmt.Println(s1)

	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()

	fmt.Println(s3)

	fmt.Println(s4)

}
