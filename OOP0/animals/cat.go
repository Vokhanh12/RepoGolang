package animals

type Cat struct {
	BaseAnimal
}

func (c Cat) Speak() string {

	return "Meow!"

}
