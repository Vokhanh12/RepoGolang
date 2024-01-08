package main

import "fmt"

func interpolationSearch(arr []int, target int) int {

	low := 0
	high := len(arr) - 1

	for low < high && target >= arr[low] && target <= arr[high] {

		probe := low + ((target-arr[low])*(high-low))/(arr[high]-arr[low])

		fmt.Printf("probe: %d\n", probe)

		if arr[probe] == target {
			return probe
		} else if arr[probe] < target {
			low = probe + 1
		} else {
			high = probe - 1
		}

	}

	return -1

}
