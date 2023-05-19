package main

import "fmt"


func main() {
	
	x := sum(1,2,3,4,5)
	fmt.Println(x)
}

func sum(values ...int)(result int){
	for _, v := range values {
		result += v
	}
	return 
}