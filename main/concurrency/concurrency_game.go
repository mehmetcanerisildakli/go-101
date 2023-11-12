package main

import (
	"fmt"
	"time"
)

func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		fmt.Println("pinger")
		time.Sleep(time.Second)
		ponger <- 1
	}
}

func ponger(pinger chan<- int, ponger <-chan int) {
	for {
		<-ponger
		fmt.Println("ponger")
		time.Sleep(time.Second)
		pinger <- 1
	}
}

func main() {
	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)
	ping <- 1
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Finished")
		return
	}
}
