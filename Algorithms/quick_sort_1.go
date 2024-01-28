package main

func quickSort_1(arr []int, start int, end int) {

	if end <= start {
		return
	}

	pivot := partition_1(arr, start, end)
	quickSort_1(arr, start, pivot-1)
	quickSort_1(arr, pivot+1, end)

}

func partition_1(arr []int, start int, end int) int {

	i := start - 1

	pivot := arr[end]

	for j := start; j <= end; j++ {

		if arr[j] < pivot {

			i++

			temp := arr[i]

			arr[i] = arr[j]

			arr[j] = temp

		}

	}

	i++

	temp := arr[i]

	arr[i] = arr[end]

	arr[end] = temp

	return i
}
