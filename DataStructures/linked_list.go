package DataStructures

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	head Node
}

func (ll *LinkedList) Add(node Node) {
	ll.addHelper(&ll.head, &node)
}

func (ll *LinkedList) addHelper(current *Node, node *Node) {

	if *current == nil {

		*current = *node

		return
	}

}

//ll.addHelper(&(*current).Next, node)

func (ll *LinkedList) PrintAll() {

	PrintAllHelper(ll.head)

}

func PrintAllHelper(node *Node) {

	if node == nil {
		return
	}

	fmt.Println(node.Data)

	PrintAllHelper(node.Next)

}
