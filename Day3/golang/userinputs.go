package main

import "fmt"

func main() {
	x := 0
	y := 0

	fmt.Print("Enter your first integer :" )
	fmt.Scanf("%d", &x )

	fmt.Print("Enter your second integer :" )
	fmt.Scanf("%d", &y )

	fmt.Println("Value of x :", x )
	fmt.Println("Value of y :", y )

	var temp string
	fmt.Println ("Press any key to exit ...")
	fmt.Scanln( &temp )
	fmt.Println( temp )

}
