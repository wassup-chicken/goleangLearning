package main

import "fmt"

func main() {
	f := incrementor()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

//func (r receiver) identifier(p parameter(s)) (return(s)) { code }
