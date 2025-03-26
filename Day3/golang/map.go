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
