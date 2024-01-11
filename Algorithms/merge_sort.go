package main

func mergeSort(arr []int) {

	length := len(arr)

	if len(arr) <= 1 {
		return
	}

	middle := length / 2

	leftArr := make([]int, middle)
	rightArr := make([]int, length-middle)

	for i := 0; i < length; i++ {
		if i < middle {
			leftArr[i] = arr[i]
		} else {
			rightArr[i-middle] = arr[i]
		}

	}

	mergeSort(leftArr)
	mergeSort(rightArr)
	merge(leftArr, rightArr, arr)

}

func merge(leftArr []int, rightArr []int, arr []int) {
	leftSize := len(leftArr)
	rightSize := len(rightArr)

	i := 0
	l := 0
	r := 0

	for l < leftSize && r < rightSize {
		if leftArr[l] < rightArr[r] {
			arr[i] = leftArr[l]
			l++
		} else {
			arr[i] = rightArr[r]
			r++
		}
		i++
	}

	for l < leftSize {
		arr[i] = leftArr[l]
		l++
		i++
	}

	for r < rightSize {
		arr[i] = rightArr[r]
		r++
		i++
	}
}
