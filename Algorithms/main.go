package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// "github.com/vokhanh12/DataStructures"

func factorial(ch chan int, quit chan int) int {

	counter := 1

	result := 1

	for {

		select {

		case ch <- counter:
			result *= counter
			counter++
		case <-quit:
			fmt.Println("quit")
			return result

		}

	}

}

func main() {
	var wg sync.WaitGroup
	// Tạo mảng ngẫu nhiên lớn hơn
	arrSize := 100
	arr := make([]int, arrSize)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}

	fmt.Println("Mảng trước sắp xếp (mảng lớn hơn):", arr)

	wg.Wait()

	fmt.Println("Mảng sau sắp xếp:", arr)

	time.Sleep(time.Second * 5)

}
