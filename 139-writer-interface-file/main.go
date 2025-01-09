package main

import (
	"log"
	"os"
)

// type person struct {
// 	first string
// }

// func (p person) writeOut(w io.Writer) error {
// 	_, err := w.Write([]byte(p.first))
// 	return err
// }

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	s := []byte("Hello World!")
	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}

}

//func (r receiver) identifier(p parameter(s)) (return(s)) { code }
