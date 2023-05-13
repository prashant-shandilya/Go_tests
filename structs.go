package main 

import (
	"fmt"
)

type person struct {
	name string,
	age int,
}

func main() {
	rahul := person{name: "Rahul", age: 20}

	fmt.Println(rahul.name)
}