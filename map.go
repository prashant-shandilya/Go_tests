package main

import (
	"fmt"
)

func main(){

	a := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
	}

	// fmt.Println(a)
	delete(a, 2)
	// ele , ok := a[2]
	fmt.Println(a[5])
	
	// fmt.Println(ele, ok)

}