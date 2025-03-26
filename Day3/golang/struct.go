package main

import "fmt"

// Creating an user-defined data-type called Rectangle with two members/attributes in it
type Rectangle struct {
	length int
	width  int
}

// This function takes no arguments and returns nothing
func sayHello() {
	fmt.Println("Hello World")
}

// go lang Method is different from function
func (rect Rectangle) Area() int {
	area := rect.length * rect.width
	return area
}

// go lang Method is different from function
func (rect Rectangle) GetLength() int {
	return rect.length
}

// go lang Method is different from function
func (rect Rectangle) GetWidth() int {
	return rect.width
}

func main() {

	rectangle := Rectangle{
		length: 100,
		width:  200,
	}

	//This demonstrates, how the function Area is associated with the struct Rectangle
	//The correct term is method as Area method can only be invoked via an instance of Rectangle data type and not directly

	fmt.Printf("Length of rectangle : %d\n", rectangle.GetLength())
	fmt.Printf("Width of rectangle : %d\n", rectangle.GetLength())
	fmt.Printf("Area of rectangle : %d\n", rectangle.Area())

	sayHello()
}
