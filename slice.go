package main

import (
	"fmt"
)

func main(){

	// build a slice and append 4 and 5 to it
	x := []int{1,2,3}
	x = append(x,4,5)
	fmt.Println(x)
	// build a slice and append another slice to it
	y := []int{6,7,8}
	x = append(x,y...)
	fmt.Println(x)

}