package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")

	/*
		numbers := make([]int, 100000)
		target := 800

		for i := 0; i < 100000; i++ {
			numbers[i] = i
		}

		startTime := time.Now()

		index := binarySearch(numbers, target)

		endTime := time.Now()
		elapsedTime := endTime.Sub(startTime)

		if index != -1 {
			fmt.Printf("Element found %d index %d\n", target, index)
		} else {
			fmt.Printf("Elememnt not found %d\n", target)
		}

		fmt.Printf("%d\n", elapsedTime)

	*/

	numbers := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}

	insertionSort(numbers)

	for _, value := range numbers {
		fmt.Print(value, " ")
	}

}
