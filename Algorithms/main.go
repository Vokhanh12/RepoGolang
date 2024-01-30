package main

import "fmt"

func main() {

	arr := []int{8, 2, 5, 3, 9, 4, 7, 6, 1}

	quickSort_3(arr, 0, len(arr)-1)

	for _, value := range arr {

		fmt.Println(value)

	}
}
