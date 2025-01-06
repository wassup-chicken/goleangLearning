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

	fmt.Println("The age of Lucas was", an["Lucas"])

	fmt.Println(an)
	fmt.Printf("%#v\n", an)
}

/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
slice:
"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "I'm 008."
Range over the records, then range over the data in each record.
*/
