package main

import "fmt"

func main() {

	//control structures
	fmt.Println("*************control structures***************")
	flag := true
	if flag {
		fmt.Println("heeey")
	} else {
		fmt.Println("hiiiiii")
	}

	for i := 7; i < 11; i++ {
		fmt.Println(i)
	}

	arr := []string{"11", "22", "33", "44"}
	for i, value := range arr {
		fmt.Println("i: ", i)
		fmt.Println(value)
	}
}
