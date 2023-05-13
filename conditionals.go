package main

import (
	"fmt"
)

func main() {
var x int  = 8
	if x > 10 || jio() || x >15{
		fmt.Println("x is greater than 10")
	}
}

func jio() bool {
	fmt.Printf("jio is called\n")

	return true
}