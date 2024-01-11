package main

import (
	"fmt"

	"github.com/vokhanh12/DataStructures"
)

func main() {

	fmt.Println("Hello world")

	ll := DataStructures.LinkedList{}

	ll.Add(&DataStructures.Node{Data: 3})
	ll.Add(&DataStructures.Node{Data: 6})
	ll.Add(&DataStructures.Node{Data: 5})
	ll.Add(&DataStructures.Node{Data: 9})

	ll.PrintAll()

}
