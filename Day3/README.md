![image](https://github.com/user-attachments/assets/6b6f5288-7c47-4778-bdd5-306a8a4b9314)# Day 3

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


## Lab - Using golang slice
Note
<pre>
- golang slice internally uses golang arrays
- golang arrays are fixed size, we can't change the size/length of the size once it is declared
- slice works like a dynamic array
- let's say initial size of slice is 5 
- when we insert 5 values into the slice it will store in temp array of size 5
- when we attempt to insert 6th value, myslice will allocate a fresh array of size 6, it copies the 5 values from the old array into the new array and then it inserts the 6th value into the newly allocated array, this is how size is able to grow in size dynamically
</pre>

Let's create a file called slice.go
<pre>
package main
import "fmt"

func main() {
	//Declares an int array of size 6 elements
	//                0   1   2   3   4   5
	intArray := [6]int{10, 20, 30, 40, 50, 60}

	fmt.Println("Array elements are ...")
	fmt.Println(intArray)

	//Slice uses an array internally
	//in this case, the slice is referring to an array from index position 2 to 4
	var mySlice []int = intArray[2:5] //2 is the lower bound index, while 5 is the upper bound of index, index 5 is not included
	fmt.Println("Slice elements are ... ")
	fmt.Println(mySlice)

	//Let's modify the slice at certain indices
	//When the slice is modified, it also affects the original array that is pointed by the slice
	mySlice[0] = 100 //mySlice[0] is nothing but intArray[2]
	mySlice[1] = 200 //mySlice[1] is nothing but intArray[3]
	mySlice[2] = 300 //mySlice[2] is nothing but intArray[4]

	fmt.Println("Slice elements after modifying slice are ...")
	fmt.Println(mySlice)

	fmt.Println("Array elements after modifying slice are ...")
	fmt.Println(intArray)

	mySlice = append(mySlice, 400) //after append, mySlice is no more pointing to intArray

	fmt.Println("Slice elements after appending new elements into slice are ...")
	fmt.Println(mySlice)

	fmt.Println("Array elements after appending new elements into slice are ...")
	fmt.Println(intArray)

	//as the intArray size is only 5, but the mySlice is attempting to add a 6th value, hence at this point slice will create a fresh array
	//mySlice and intArray are not associated from this point onwards
	mySlice = append(mySlice, 500) //after append, mySlice is no more pointing to intArray

	fmt.Println("Slice elements after appending new elements into slice are ...")
	fmt.Println(mySlice)

	fmt.Println("Array elements after appending new elements into slice are ...")
	fmt.Println(intArray)
}
</pre>

Run it
```
go run ./slice.go
```

Expected output

![image](https://github.com/user-attachments/assets/60d9875e-ea50-47cf-b524-c24e33ef80ce)
![image](https://github.com/user-attachments/assets/34edc6c0-fe9a-492d-99a6-6e5f359f9527)

## Lab - Golang if-else control statement
Create a file named if-else.go

```
package main

import "fmt"

func main() {

	x := 100

	if x%2 == 0 {
		fmt.Println(x, " is an even number")
	} else {
		fmt.Println(x, " is an on odd number")
	}

}
```

Run it
```
go run ./if-else.go
```

Expected output
![image](https://github.com/user-attachments/assets/ecbd9956-893f-4851-a433-0fc0a88e80c5)

## Lab - Golang map
Create a file named map.go
```

package main

import "fmt"

func main() {

	//Declares a map with key of type string declared within square bracket, value of type string followed by square bracket end
	//key and values can be of different data types
	toolsPath := map[string]string{
		"java_home": "/usr/lib/jdk-11",
		"mvn_home":  "/usr/share/maven",
	}

	fmt.Println("Java Home Directory : ", toolsPath["java_home"])

	//add a key,value pair into an existing map
	toolsPath["go_home"] = "/usr/go"

	//iterating a map and printing its values
	for key, value := range toolsPath {
		fmt.Println(key, value)
	}

	//delete a key-vaule pair from an existing map
	delete(toolsPath, "go_home")
	fmt.Println(toolsPath)

}
```

Run it
```
go run ./map.go
```

Expected output
![image](https://github.com/user-attachments/assets/646b6af1-5e64-495f-8f50-0de844872bae)

## Lab - Using struct in go lang

Create a file named struct.go
```
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

func main() {
	rectangle := Rectangle{
		length: 100,
		width:  200,
	}

	//This demonstrates, how the function Area is associated with the struct Rectangle
	//The correct term is method as Area method can only be invoked via an instance of Rectangle data type and not directly
	fmt.Printf("Area of rectangle : %d\n", rectangle.Area())

	sayHello()
}
```

Run it
```
go run ./struct.go
```

Expected output
![image](https://github.com/user-attachments/assets/2c2f19f5-ad9f-4e3f-8a3c-3d0486765bab)
![image](https://github.com/user-attachments/assets/86298bb6-a284-4f63-be2a-dce0d57526e9)
![image](https://github.com/user-attachments/assets/264d2987-e29e-49d2-a631-011eb9101f0d)

## Info - Golang modules
<pre>
- is a collection one or more packages
- you could write reusable code
- there can multiple versions of the same module as well
</pre>

## Lab - Creating Golang modules

Let's create two modules namely addition and subtraction, hence let's create two folder 
```
cd ~
mkdir addition subtraction
```

Let's create a module named addition
```
cd addition
go mod init addition  //Creates a file called go.mod with the name of the module and go language version supported
```

Under the addition folder, let's create a file called add.go with the below content
<pre>
package addition
//The functions exposed for external access must begin with upper case 
func Add( x float32, y float32 ) float64 {
	return float64(x + y)
}
</pre>

Under the subtraction folder
```
cd ~/subtraction
go mod init subtraction //Creates a file named go.mod with the name of the module and go language version supported
```

Under the subtraction folder, let's create a file called subtraction.go with the below content
<pre>
package subtraction

//The functions exposed for external access must begin with upper case 
func Subtract( x float32, y float32 ) float64 {
	return float64(x-y)
}
</pre>

Let's create a third module called main
```
cd ~/golang-practice
go mod init main
```

Let's create a file named main.go with the below code
<pre>
package main

import (
	"fmt"
	"addition"
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
</pre>

Run it
```
go mod tidy  //This is going to download all the dependent modules
go mod edit --replace addition=../addition  // This helps golang to locate the addition module
go mod edit --replace subtraction=../subtraction //This helps golang to locate the subtraction module

go run ./main.go
```

Expected output
![image](https://github.com/user-attachments/assets/e3a00b4c-01ca-4506-b930-2061f5d68445)
![image](https://github.com/user-attachments/assets/6e511456-b1e8-4b52-b046-5c6c1aa6c9ff)

Module Versioning
<pre>
- In case you have multiple versions of the same module, you need to create a subfolder for each version and copy the go.mod and all your *.go files into it
</pre>

Let's create a second version for module addition
```
cd ~/addition
mkdir v2
cp * v2
cat go.mod
cat add.go
cd ..
go mod edit --replace addition/v2=./addition/v2
cat go.mod
cat main.go
go run ./main.go
```


Expected output
![image](https://github.com/user-attachments/assets/6bf0d990-ec75-47b4-a10e-2adaf9d25c65)
![image](https://github.com/user-attachments/assets/87447613-96a3-4ad7-a765-5f36254e4547)
![image](https://github.com/user-attachments/assets/beae3aca-b0cb-4aa4-8480-d5f2bbda0ff0)

## Lab - Building the golang application into a binary
```
cd golang/module
cat go.mod
go build .
./main
```

Expected output
![image](https://github.com/user-attachments/assets/44473afc-386f-417d-a819-cebd481d8450)


## Lab - Golang concurrent functions

Create a file named main.go
```
package main

import (
	"fmt"
	"time"
)

func firstFunction( count int ) { 
  for i := 0; i<count; i++ {
    fmt.Println("First function " , i )
    time.Sleep( time.Millisecond * 5 )
  }
}

func secondFunction( count int ) { 
  for i := 0; i<count; i++ {
    fmt.Println("Second function " , i )
    time.Sleep( time.Millisecond * 5 )
  }
}

func main() {
   fmt.Println ("Press any key to exit ...")

   //Invoking them in sequence, once the firstFunction completes running then only the secondFunction will start running
   firstFunction(1000)
   secondFunction(1000)

   //We wish to run both firstFunction and secondFunction in parallel ( at the same time )
   go firstFunction(1000)
   go secondFunction(1000)

   var tmp string
   fmt.Scanln(&tmp) //this will make sure the program waits until some key is pressed to exit

}
```

Let run it
```
go run ./main.go
```

Expected output
![image](https://github.com/user-attachments/assets/0fcab44f-b0ac-4a47-bc58-c6138519f010)
![image](https://github.com/user-attachments/assets/2d39736b-d516-4e94-94e6-12e09f7d2ae5)
![image](https://github.com/user-attachments/assets/c72c183a-19c1-4f5b-81a1-8386b9936d2b)

## Info - Terraform Overview
<pre>
- is a Infrastructure as a code tool
- it is used to automate provisioning containers, pull container images, virtual machine locally or on public cloud etc.,
- it can also be used to provision storage locally(your private datacenter), or in public/hybrid cloud
- it can help provisioning virtual machines using virtualbox, vmware, vsphere, etc.,
- Terraform is cloud newtral, it works pretty much in any enviroment and all the private, public and hybrid clouds
- Alternates to Terraform
  - AWS Cloudformation
    - disadvantage - can be used only to automate infrastructure provision in AWS public cloud
- is developed in go language by a company called HashiCorp
- the Domain Specific Language(DSL) i.e the language in which the autotomation code is written
- Terraform uses HCL(Hashicorp Configuration Language) as the the DSL to automate infrastructure declaratively
- it comes in two flavours
  1. Terraform Core - open source and supports only command line
  2. Terraform Enterprise - requires license and supports Web Interface, all world-wide support is provided by HashiCorp
</pre>

## Info - Terraform High Level Architecture
![Terraform](terraform-architecture-diagram.png)

## Info - Terraform Providers
<pre>
- Terraform depends on Providers to provision resources
- For example
  - In order to provision an ec2 instance in AWS, Terraform depends on a provider called AWS
  - In order to provision an azure VM, in Azure portal, Terraform depends on a provider called Azure
  - as long as there is a provider, Terraform can provision such resources
- In case, to provision a particular type of Resource within your organization there is no ready-made provider, you can develop your own Provider in Golang using Terraform Provider SDK
- Providers supports two types of objects
  1. Resources
     - If you wish to Provision ec2 using Terraform, then you will define a resource block expressing your expected state in declarative style
     - Terraform can create, Replace, Update, Delete the resources managed by Terraform
  2. Datasources
	- In case to provision ( create certain resources ), your declarative terraform script(manifest) file depends on already existing resources, then we call them as Datasources or Data block
	- These are read-only resources, Terraform can use it but can't modify or delete it
</pre>

## Info - Terraform Resources
<pre>
- Each Provider supports one to many Resources
- For instance the docker provider supports the following resources
  - docker_image
  - docker_container
</pre>

## Info - Terraform DataSources
<pre>
- Each Provider supports zero to many Datasources
- For instance, the docker provider supports the following Datasources (Read only resources)
  - docker_image
  - docker_container
</pre>

## Lab - Checking the Terraform version
```
terraform version
```

Expected output
![image](https://github.com/user-attachments/assets/0e38596a-0bdf-4266-8c1e-1da2e7e59e46)

## Info - Terraform Providers can be download from registry.terraform.io website
![image](https://github.com/user-attachments/assets/82f246f4-1397-4dc5-96fd-7cc6c197fba6)

## Lab - Using Docker Image to provision containers
Create a file named main.tf
```
terraform {
  required_providers {
    docker = {
      source = "kreuzwerker/docker"
      version = "3.0.2"
    }
  }
}

# This is an instance of the provider with optional configurations
provider "docker" {
  # Configuration options
}

data "docker_image" "tektutor_docker_image" {
   name = "tektutor/ubuntu-ansible-node:latest"
}
```

In order to download the terrform docker provider from registry.terraform.io website, we need run the below command
```
terraform init
tree .terraform
cat .terraform.lock.hcl
cat main.tf
terraform plan
terraform apply
cat terraform.tfstate
docker ps
```

Expected output
![image](https://github.com/user-attachments/assets/e170c6a4-ce18-4834-bf08-f594677bebff)
![image](https://github.com/user-attachments/assets/d1ae5d96-bc83-4645-b56d-8af5948a3092)
![image](https://github.com/user-attachments/assets/f8d8c5e2-815d-46a7-b18e-3dc476e87f6f)
![image](https://github.com/user-attachments/assets/b0dc4a0e-b7a5-4635-9341-16c34b3f3308)
![image](https://github.com/user-attachments/assets/9e6473d2-6e50-406d-a028-23a44536a8be)
![image](https://github.com/user-attachments/assets/d4213785-115a-45f6-a93b-986796a4e9af)
![image](https://github.com/user-attachments/assets/ab8c8930-775a-4d96-962c-54c615b09a83)
![image](https://github.com/user-attachments/assets/824397d6-45d7-44ea-8379-72ba3058dd33)
![image](https://github.com/user-attachments/assets/9103874f-a0a3-48fb-8458-8fb8d47ecd4c)
![image](https://github.com/user-attachments/assets/e122cb04-b878-4774-980f-6df47092e7dc)
![image](https://github.com/user-attachments/assets/e9e0b660-9316-4537-be7a-f6175238f9a6)
![image](https://github.com/user-attachments/assets/f6011e51-b74a-4925-878e-fea004f02d2b)
![image](https://github.com/user-attachments/assets/f5b5b811-f65d-4d0f-8eba-28cc1ae97b87)
![image](https://github.com/user-attachments/assets/6e536b8e-8587-4e0d-bc74-148115dd5573)
![image](https://github.com/user-attachments/assets/3dd30813-a4b6-41e8-b25f-df41937d06d5)
![image](https://github.com/user-attachments/assets/12878d9b-b95b-4b5c-9d2b-5835eb96625b)
![image](https://github.com/user-attachments/assets/654a3317-f459-4bcf-b1d1-2f4b2e3fc683)
![image](https://github.com/user-attachments/assets/0a0120ff-947c-40cc-948f-a5637e9b8aeb)
![image](https://github.com/user-attachments/assets/604895dd-ed10-47ab-849f-0c3b4c20bd6d)


