package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting the program")
	defer fmt.Println("This line will never be called")
	defer fmt.Println("Sigma Man")
	panic("A severe error occurred: stopping the program!")
	
	fmt.Println("Ending the program")

}