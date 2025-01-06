package main

import "fmt"

func main() {
	am := map[string]int{
		"Hong":  42,
		"Seung": 32,
		"Pam":   28,
	}

	fmt.Println("The age of Hong was", am["Hong"])

	fmt.Println(am)
	fmt.Printf("%#v\n", am)

	an := make(map[string]int)
	an["Lucas"] = 31
	an["Stephanie"] = 37
	an["George"] = 78

	fmt.Println("The age of Lucas was", an["Lucas"])
	fmt.Println(an)
	fmt.Printf("%#v\n", an)

	delete(an, "George")

	fmt.Println(an["George"])

	v, ok := an["Lucas"]
	if ok {
		fmt.Println("it prints!!!", v)
	} else {
		fmt.Println("Key didn't exist")
	}

	// for range over a MAP
	for k, v := range an {
		fmt.Println(k, v)
	}
	//just value
	for _, v := range an {
		fmt.Println(v)
	}
	//just key
	for k := range an {
		fmt.Println(k)
	}

	// for rannge with a SLICE

	xi := []int{42, 43, 44}

	for i, v := range xi {
		fmt.Println(i, v)
	}
	//just value
	for _, v := range xi {
		fmt.Println(v)
	}
	//just index
	for i := range xi {
		fmt.Println(i)
	}
}

/*
test123
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
slice:
"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "I'm 008."
Range over the records, then range over the data in each record.
*/
