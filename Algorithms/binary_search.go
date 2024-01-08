package main

import "fmt"

func binarySearch(arr []int, target int) int {

	low := 0

	high := len(arr) - 1

	for low <= high {

		middle := (low + high) / 2

		fmt.Printf("middle: %d", middle)

		if target < arr[middle] {
			high = middle - 1
		} else if target > arr[middle] {
			low = middle + 1
		} else if target == arr[middle] {
			return target
		}

	}

	return -1

}
