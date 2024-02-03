package main

func quickSort_4(arr []int, start int, end int) {

	if end <= start {
		return
	}

	pivot := partition_4(arr, start, end)

	quickSort_4(arr, start, pivot-1)
	quickSort_4(arr, pivot+1, end)

}

func partition_4(arr []int, start int, end int) int {

	pivot := arr[end]

	i := start - 1

	for j := start; j <= end; j++ {

		if arr[j] < pivot {

			i++

			temp := arr[j]
			arr[j] = arr[i]
			arr[i] = temp

		}
	}

	i++

	temp := arr[i]
	arr[i] = arr[end]
	arr[end] = temp

	return i

}
