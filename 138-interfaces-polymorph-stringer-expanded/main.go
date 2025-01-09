package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("this is the book ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("the number is ", strconv.Itoa(int(c)))
}

func logInfo(s fmt.Stringer) {
	log.Println("LOG FROM 138", s.String())
}

func main() {
	b := book{
		title: "West With The Night",
	}

	var c count = 42
	logInfo(b)
	log.Println(c)
}

//func (r receiver) identifier(p parameter(s)) (return(s)) { code }
