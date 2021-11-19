package main

import (
	"fmt"
	"time"
)

func example() {
	tick := time.Tick(time.Second)
	after := time.After(3 * time.Second)
	channel := make(chan int, 1)
	go func() {
		channel <- 1
		fmt.Println("======= goroutine ======")
		close(channel)
	}()
	for {
		select {
		case <-tick:
			fmt.Println("tick 1 second")
		case <-after:
			fmt.Println("after 3 second")
			return
		case value, ok := <- channel:
			if ok {
				fmt.Println("got", value)
			} else {
				fmt.Println("channel is closed")
				time.Sleep(time.Second)
			}
		default:
			fmt.Println("come into default")
			time.Sleep(500 * time.Millisecond)
		}
	}
}


func main() {
	example()
}
