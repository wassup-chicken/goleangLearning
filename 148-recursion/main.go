package main

import "fmt"

func main() {
	x := factorial(4)
	fmt.Println(x)

	y := factLoop(4)
	fmt.Println(y)
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func factLoop(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--
	}

	return total
}
