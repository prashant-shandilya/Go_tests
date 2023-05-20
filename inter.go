package main

import (
	"fmt"
)

func main() {

	x := []byte("Hello Guys")
	fmt.Println(x)
}

type writer interface {
	Write([]byte) (int, error)
}
