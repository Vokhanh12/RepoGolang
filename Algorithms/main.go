package main

import (
	"fmt"
	"time"
)

// "github.com/vokhanh12/DataStructures"

func main() {

	var c = make(chan int, 4)

	go func() { fmt.Printf("got %d\n", <-c) }()

	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 4

	time.Sleep(time.Second)

}
