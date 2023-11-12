package main

import (
	"fmt"
	"time"
)

func heavy() {
	for {
		time.Sleep(time.Second)
		fmt.Println("waiting ^^^")
	}
}

func main() {

	//concurrency
	//fmt.Println("*************concurrency***************")

	//fmt.Println("started")
	//heavy()
	//go heavy()
	//time.Sleep(time.Second * 3) // we give the time to run heavy() or you can use select{}
	//fmt.Println("finished")

	//wait
	//fmt.Println("*************wait***************")
	//wg := &sync.WaitGroup{}
	//wg.Add(2)
	//
	//go func() {
	//	fmt.Println("xxxxxxxx")
	//	wg.Done()
	//}()
	//go func() {
	//	fmt.Println("qqqqqqqqq")
	//	wg.Done()
	//}()
	////time.Sleep(time.Second * 3)
	//wg.Wait()

	//channels
	fmt.Println("*************channels***************")
	c := make(chan int)
	fmt.Println(c)

	go func() {
		c <- 1
	}()

	val := <-c
	fmt.Println(val)
	fmt.Println(c)

	c2 := make(chan int, 3)

	go func() {
		c2 <- 1
		c2 <- 2
		c2 <- 3
		c2 <- 4
		close(c2)
	}()

	for i := range c2 {
		fmt.Println(i)
	}

}
