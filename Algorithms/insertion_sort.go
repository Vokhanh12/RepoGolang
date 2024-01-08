package main

// 9, 1, 8, 2, 7, 3, 6, 4,
// 1, 8, 9
func insertionSort(arr []int) {

	for i := 1; i < len(arr); i++ {

		temp := arr[i]

		for j := i - 1; j >= 0; j-- {

			if temp < arr[j] {

				arr[j+1] = arr[j]

			} else {

				arr[j] = temp

			}

			if j == 0 {
				arr[j] = temp

			}

		}

	}

}
