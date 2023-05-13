package main 

import (
	"fmt"
)



func main() {

	type person struct {
		name string
		age int
	}

	rahul := person{name: "Rahul", age: 20}

	fmt.Println(rahul.age)
}