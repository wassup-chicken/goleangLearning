package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	xb, err := readFile("poem.txt")
	if err != nil {
		log.Fatalf("readFile in main: %s", err)

	}
	fmt.Println(xb)
	fmt.Println(string(xb))
}

// wrapper function
func readFile(fileName string) ([]byte, error) {
	//wrapped function
	xb, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error in readFile func: %s", err)
	}
	return xb, nil
}
