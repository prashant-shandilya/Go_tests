package main 

import (
	"fmt"
)

func main() {
	// make a struct with 3 fields and print it
	type person struct {
		first string
		last string
		age int
	}
	p1 := person{"James", "Bond", 34}
	fmt.Println(p1)
}