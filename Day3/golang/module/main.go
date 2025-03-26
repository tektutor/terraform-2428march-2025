package main

import (
	"fmt"
	"addition/v2"
	"subtraction"
)

func main() {

	x := float32(100.123) //We are casting/converting float64 into float32
	y := float32(200.456) //We are casting/converting float64 into float32
	//x := 100.123  By default go lang will assume 100.123 as float64, hence we will get compilation error
	//y := 200.456  By default go lang will assume 200.456 as float64, hence we will get compilation error

	fmt.Println ( "The sum of ", x, " and ", y, " is ", addition.Add( x, y ) )
	fmt.Println ( "The difference of ", x, " and ", y, " is ", subtraction.Subtract( x, y ) )

}
