package main

import "fmt"

// Rectangle is a struct width 3 fields
// structs are typed collections of fields
type Rectangle struct {
	Width       int
	Height      int
	Description string
}

func main() {

	aRectangle := Rectangle{
		Width:       10,
		Height:      2,
		Description: "just a simple rectangle",
	}

	fmt.Println("An empty Rectangle struct (omitted fields are zero-valued): ", Rectangle{})
	fmt.Println("Our Rectangle struct: ", aRectangle)

	fmt.Println("We can access the fields with a dot:")
	fmt.Println("Rectangle height:", aRectangle.Height)
	fmt.Println("Rectangle weight:", aRectangle.Width)

	aRectangle.Height = 5
	fmt.Println("We can assign values to struct fields with the = sign")
	fmt.Println("Rectangle height: ", aRectangle.Height)

	rectanglePointer := &Rectangle{
		Width:  3,
		Height: 3,
	}

	fmt.Println("Using the & sign, will yield a pointer to a Rectangle struct: ", rectanglePointer)
	fmt.Println("We can dereference by using the * symbol:", *rectanglePointer)

	fmt.Println("You can change the address that your pointer points to as well ")
	fmt.Println("My old address: ", fmt.Sprintf("%p", rectanglePointer))
	rectanglePointer = &aRectangle
	fmt.Println("My new address: ", fmt.Sprintf("%p", rectanglePointer))
}
