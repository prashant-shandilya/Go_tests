package main

import ( 
	"fmt"
 )


func main() {
	f := func (x int){
		fmt.Println("Hello, ", x)
	}

	f(5)

	}