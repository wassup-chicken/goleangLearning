package main

import "fmt"

type person struct {
	first string
	last  string
	favIC []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		favIC: []string{"choco", "banana", "passionfruit"},
	}

	p2 := person{
		first: "Hi",
		last:  "World",
		favIC: []string{"choco", "strawberry", "vanilla"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v)
		for _, v2 := range v.favIC {
			fmt.Println(v.first, v.last, v2)
		}
	}

}
