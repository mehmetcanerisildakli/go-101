package main

import (
	"fmt"
	"time"
)

func selected(c chan int, quits chan struct{}) {
	for {
		select {
		case <-c:
			fmt.Println("received")
		case <-quits:
			fmt.Println("quit")
			break
		}
	}
}

func main() {
	c := make(chan int)
	quits := make(chan struct{})
	go func() {
		c <- 1
	}()
	go selected(c, quits)

	time.Sleep(time.Second * 5)
}
