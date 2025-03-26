![image](https://github.com/user-attachments/assets/d8e57b40-c132-464e-aaf1-65a030468dca)# Day 3

## Lab - In case you prefer using Visual Studio Code Editor for Golang application development
You can install Visual Studio Code editor in Ubuntu as shown below, when prompts for password type rps
```
sudo snap install code --classic
```

## Lab - Writing your first Hello World application in Golang
Create a file named hello.go with below content

```
package main

import "fmt"

func main() {
   fmt.Println ( "Hello World!" )
}
```

Note
<pre>
- main is the entry-point function in go language i.e the default function that will be invoked first you run any go application
- every go application must have one main function under main package
- fmt package has input/output functions like 
  - Print
  - Println
  - Sprint
  - Sprintf
  - Sprintln
  - Fprint,
  - Fprintf
  - Fprintln
  - Scan
  - Scanf
  - Scanln
  - Sscan
  - Sscanln
  - Fscan
  - Fscan
  - Fscanln
  - Errorf
</pre>

To run the application, you can issue the below command
```
mkdir ~/golang-practice
cd golang-practice
cat hello.go
go run ./hello.go
```
Expected output
![image](https://github.com/user-attachments/assets/d181a189-1ec0-4899-8eb3-868e4a0dd074)
![image](https://github.com/user-attachments/assets/3cfe284a-8ed8-415d-90f4-cdfd090d41dc)

## Lab - Go lang loop
<pre>
- Go language supports only for loop  
</pre>

You can create a file named loops.go with the below content
<pre>
package main

import "fmt"

func main() {

	count := 5 //Declares a count variable of type int and assigns a value 5

	for count > 0 {
		fmt.Println(count)

		count-- //equivalent to count = count - 1
		//--count  pre-decrement is not supported in golang unlike C++
		//++count  pre-increment is not supported in golang unlike C++
		//count++ is supported in golang
	}

	fmt.Println("Value of count is", count, "after for loop")

	count = 0 //Variable is already declared, in this line we are just resetting the count value to 0

	for count = 1; count < 10; count++ {
		fmt.Printf("%d\t", count)
	}

	fmt.Println()

}   
</pre>

Run it
```
go run ./loops.go
```

Expected output
![image](https://github.com/user-attachments/assets/271bb216-7785-4418-9690-f2c345e8b1e2)

## Lab - Using Switch case in Go language
Create a file named switch-case.go with the below code
```
package main

import "fmt"

func main() {

	var direction string //declares a variable named direction of type string

	fmt.Print("Enter some direction :")
	fmt.Scanln(&direction)

	switch direction {

	case "east":
		fmt.Println("You entered direction ", direction)
	case "west":
		fmt.Println("You entered direction ", direction)
	case "north":
		fmt.Println("You entered direction ", direction)
	case "south":
		fmt.Println("You entered direction ", direction)
	default:
		fmt.Println("Invalid direction, possible values are east, west, north, south")
	}
}
```
Run it
```
go run ./switch-case.go
```

Expected output
![image](https://github.com/user-attachments/assets/e9149048-aace-4e73-870b-301d77730a37)

## Lab - Go lang arrays

Create a file named arrays.go with the below content
```
package main

import "fmt"

func main() {

	//We have declared an array of integers of size 5
	//So we can store upto 5 integer values into this array
	//go lang array size is fixed
	//array index starts from 0
	//valid array index range is 0 to 4, total 5 values
	var arr [5]int

	//let's assign some values into the array
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	//arr[5] = 60 This will report array index out of bounds error

	fmt.Println("Array elements are ...")
	fmt.Println(arr)

	count := len(arr)
	fmt.Println("Length of array :", count)

	//Modifying values stored in an array
	arr[3] = 25

	fmt.Println("Array elements are ...")
	for i := 0; i < count; i++ {
		fmt.Printf("%d\t", arr[i])
	}
	fmt.Println()
}
```

Run it
```
go run ./arrays.go
```

Expected output
![image](https://github.com/user-attachments/assets/43b6969e-0bf7-4dd2-bb09-3e58712af9e3)
![image](https://github.com/user-attachments/assets/765dafe2-b5c8-4c12-ab58-c8dc21aad377)

## Lab - Go lang user-defined functions

Create a file named functions.go with the below content
```
package main

import "fmt"

func main() {
	fmt.Println(sayHello("Golang"))
}

// This function accepts a string input and returns a string output
func sayHello(msg string) string {
	return "Hello " + msg + " !"
}
```

Run it
```
go run ./functions.go
```

Expected output
![image](https://github.com/user-attachments/assets/b2b5f097-67e1-4b83-8201-5431c4d67567)

## Lab - Functions with multiple return values

Create a file named function-with-multiple-returns.go
```
package main

import "fmt"

// myFunction takes no parameters but return two integers
func myFunction() (int, int) {
	return 10, 20
}

func main() {

	x, y := myFunction() // := is short form of declaring a new variable and initializing a value

	fmt.Println("Value of x is ", x)
	fmt.Println("Value of y is ", y)

}

```

Run it
```
go run ./function-with-multiple-returns.go
```

Expected output

![image](https://github.com/user-attachments/assets/55fcf524-c68e-420e-ac8b-da9d510eb747)

## Lab - Golang pointer - pass by pointer
Create a file named pointers.go
```
package main

import "fmt"

// This function takes a string pointer as an input parameter
func sayHello(msgPtr *string) {

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
```

Run it 
```
go run ./pointers.go
```

Expected output
![image](https://github.com/user-attachments/assets/b67c1a5d-d5b0-4c1f-9760-26e03ff98ea4)
