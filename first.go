package main

import (
	"fmt"
)

func main() {

	x := map[string]int{
		"Sydney":43,
		"Canberra":32,
	}
	x["Delhi"] = 88
	delete (x,"Sydney")
	fmt.Println(len(x))
}