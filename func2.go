package main

import (
	"fmt"
)

func main() {
	res , err:= divide(10, 0)
	if err != nil{
		fmt.Println(err)
		return 
	}
	fmt.Println(res)
}
func divide(a,b int) (int,error){
	if b==0{
		return 0,fmt.Errorf("Cannot divide by zero")
	}
	return a/b,nil
}