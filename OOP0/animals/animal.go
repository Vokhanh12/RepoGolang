package animals

type Animal interface {
	Speak() string
	GetName() string
}

type BaseAnimal struct {
	Name string
	Age  string
}

func (a BaseAnimal) Speak() string {

	return "Generic animal sound"

}

func (a BaseAnimal) GetName() string {
	return a.Name
}
