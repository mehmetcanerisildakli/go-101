package main

import "fmt"

var var2 = 11
var var3 int

func main() {
	n, err := fmt.Println("h1", 123, false, 'a', var2, var3)

	_, e := fmt.Println(n, err)
	if e != nil {
		fmt.Println("error", e)
	}
}
