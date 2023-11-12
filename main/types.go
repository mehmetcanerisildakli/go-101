package main

import (
	"fmt"
)

var a int

type pizza int

var b pizza

func main() {
	a = 11
	b = 22
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	a = int(b) // conversion
	fmt.Println(a)
}
