# Day 3

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
