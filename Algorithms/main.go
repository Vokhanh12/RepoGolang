package main

import "fmt"

//"github.com/vokhanh12/DataStructures"

func main() {
	/*

		fmt.Println("Hello world")

		ll := DataStructures.LinkedList{}

		ll.Add(DataStructures.Node{Data: 3})
		ll.Add(DataStructures.Node{Data: 6})
		ll.Add(DataStructures.Node{Data: 5})
		ll.Add(DataStructures.Node{Data: 9})
		ll.Add(DataStructures.Node{Data: 8})
		ll.Add(DataStructures.Node{Data: 7})

		ll.PrintAll()

		fmt.Println("SumEven:", ll.SumEven())

	*/

	numbers := []int{3, 7, 8, 5, 4, 2, 6, 1}

	mergeSort(numbers)

	for _, value := range numbers {
		fmt.Println("", value)
	}

}
