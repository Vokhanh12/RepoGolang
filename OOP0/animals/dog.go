package animals

type Dog struct {
	BaseAnimal
}

func (d Dog) Speak() string {

	return "Woof!"

}
