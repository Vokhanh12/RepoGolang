package main

import (
	"fmt"

	"github.com/vokhanh12/hellogo/animals"
)

func main() {

	cat := animals.Cat{BaseAnimal: animals.BaseAnimal{Name: "jack", Age: "3"}}

	dog := animals.Dog{BaseAnimal: animals.BaseAnimal{Name: "lucy", Age: "2"}}

	fmt.Println(cat.Speak())

	fmt.Println(dog.Speak())

}
