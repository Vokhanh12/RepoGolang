package main

func partition(array []int, start int, end int) int {

	pivot := array[end]

	i := start - 1
	j := start

	for j <= end {

		if pivot < array[j] {

			j++

			if j == end {

				temp := array[i]

				array[i] = array[j]

				array[j] = temp

			}

		} else {

			i++

			temp := array[i]

			array[i] = array[j]

			array[j] = temp

		}

	}

	return i

}

func quickSort(array []int, start int, end int) {

	if start >= end {
		return
	}

	pivot := partition(array, start, end)

	quickSort(array, start, pivot-1)
	quickSort(array, pivot+1, end)

}
