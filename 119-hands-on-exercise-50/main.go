package main

import "fmt"

func main() {
	m := make(map[string][]string)
	m[`bond_james`] = []string{`shaen, not stirred`, `martinis`, `fast cars`}
	m[`moneypenny_jenny`] = []string{`shaen, not stirred`, `martinis`, `fast cars`}
	m[`no_dr`] = []string{`cats,`, `ice cream`, `fast cars`}

	fmt.Printf("%#v\n", m)

	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}
}
