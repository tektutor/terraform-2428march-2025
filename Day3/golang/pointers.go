package main

import "fmt"

// This function takes a string pointer as an input parameter
func sayHello(msgPtr *string) {

	//Dereferrencing means, the values stored at address pointed by msgPtr will be printed here
	fmt.Println("Inside sayHello ", *msgPtr)
	fmt.Println("Address pointed by msgPtr is ", msgPtr)

	tmp := *msgPtr //The value stored at the address pointed by msgPtr is assigned to tmp string

	//The value stored at address pointed by msgPtr we are changing to "Hello Golang !"
	*msgPtr = tmp + " Golang" + " !"

	fmt.Println("Inside sayHello before return ", *msgPtr)

}

func main() {
	//msg is a string variable assigned with "Hello" value
	msg := "Hello"
	fmt.Println("Message before calling sayHello function is ", msg)
	fmt.Println("Address of msg string is ", &msg)
	//SayHello function is taking the address of msg string
	sayHello(&msg)
	fmt.Println("Message after calling sayHello function is ", msg)
}
