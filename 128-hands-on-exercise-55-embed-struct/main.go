package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {

	v1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Ford",
		model: "Mustang",
		doors: 4,
		color: "yellow",
	}
	v2 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Toyota",
		model: "Tundra",
		doors: 4,
		color: "White",
	}
	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println(v1.electric, v1.make, v1.model)
	fmt.Println(v2.electric, v2.make, v2.model)
}
