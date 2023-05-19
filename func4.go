package main

import (
	"fmt"
)

type college struct {
	name string
	year int
}

func main() {
	g := college{"IIT", 1950}
	g.say()
}
func (c college) say() {
	fmt.Println("Hello, ", c.name)
} 