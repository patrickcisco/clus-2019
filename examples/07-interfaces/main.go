package main

import "fmt"

type Dog struct {
	Type string
	Name string
}

type Cat struct {
	Type string
	Name string
}

func (v *Dog) Speak() string {
	return fmt.Sprintf("woof, I'm a %v", v.Type)
}

func (v *Cat) Speak() string {
	return fmt.Sprintf("meow, my name is %v", v.Name)
}

type Animal interface {
	Speak() string
}

func makeNoise(a Animal) {
	sound := a.Speak()
	fmt.Println(sound)
}

func main() {
	dog := &Dog{
		Name: "Baxter",
		Type: "Border Terrier",
	}

	cat := &Cat{
		Type: "Tabby",
		Name: "Garfield",
	}

	makeNoise(cat)
	makeNoise(dog)

}
