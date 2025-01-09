package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())
	b.WriteString("Gophers!")
	fmt.Println(b.String())
	b.Reset()
	b.WriteString("Tis Thursday")
	fmt.Println(b.String())

}

//func (r receiver) identifier(p parameter(s)) (return(s)) { code }
