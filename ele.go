package main

import (
	"fmt"
)

func ele()  bool {

defer fmt.Println("deferred statement executed")
return true;
}

func main() {
	fmt.Printf("I am in main function\n")
ele()

x := []byte("mohammed ali")
fmt.Println(x)

}

