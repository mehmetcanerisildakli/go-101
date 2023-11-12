package main

import "fmt"

func anything(a interface{}) {
	fmt.Println(a)
}
func main() {
	anything(3.14)
	anything("vvv")
	anything(struct {
	}{})

	mymap := make(map[string]interface{})
	mymap["caner"] = "xo4562029874nbvhy"
	fmt.Println(mymap)
}
