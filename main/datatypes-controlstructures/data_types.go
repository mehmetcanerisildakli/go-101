package main

import "fmt"

type MyCar struct {
	Brand string
	Model string
	year  int
}

func main() {
	// numbers
	fmt.Println("*************Numbers***************")
	var (
		m1 = 234
		m2 = 465
	)
	m3 := 987 // we can use this while declaration a variable
	fmt.Println(m1 + m2 + m3)

	// strings
	fmt.Println("*************Strings***************")

	m4 := " trendyol "
	var m5 = "go"
	fmt.Println(m5 + m4 + m5)

	// arrays
	fmt.Println("*************Arrays***************")

	var arr []int
	arr = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)
	arr = append(arr, 7, 8, 9)
	fmt.Println(arr)

	//pointers
	fmt.Println("*************Pointers***************")

	p1 := 3
	p2 := &p1
	fmt.Println(p2)
	fmt.Println(*p2)

	//structures
	fmt.Println("*************Structures***************")

	c1 := MyCar{"Tesla", "Model 3", 2024}
	fmt.Println(c1)
}
