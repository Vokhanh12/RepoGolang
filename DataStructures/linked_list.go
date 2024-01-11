package DataStructures

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Add(node Node) {
	ll.addHelper(&ll.head, node)
}

func (ll *LinkedList) addHelper(current **Node, node Node) {

	if *current == nil {

		*current = &node

		return
	}

	ll.addHelper(&(*current).Next, node)

}

func (ll *LinkedList) PrintAll() {

	printAllHelper(ll.head)

}

func printAllHelper(node *Node) {

	if node == nil {
		return
	}

	fmt.Println(node.Data)

	printAllHelper(node.Next)

}

func (ll *LinkedList) SumEven() int {

	return sumEvenHelper(&ll.head)

}

func sumEvenHelper(node **Node) int {

	if *node == nil {
		return 0
	}

	if (*node).Data%2 == 0 {
		return sumEvenHelper(&(*node).Next) + 1
	} else {
		return sumEvenHelper(&(*node).Next) + 0
	}

}
