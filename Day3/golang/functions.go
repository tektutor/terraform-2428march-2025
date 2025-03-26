package main

import "fmt"

func main() {
	fmt.Println(sayHello("Golang"))
}

// This function accepts a string input and returns a string output
func sayHello(msg string) string {
	return "Hello " + msg + " !"
}
