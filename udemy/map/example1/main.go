package main

import "fmt"

func main() {
	//MAP is defined as key and value both are of type string
	var Greeting = make(map[string]string)

	//Composite literal
	//var Greeting = map[string]string{}

	Greeting["virendra"] = "Ashima"
	Greeting["Aarav"] = "Bisht"

	fmt.Println(Greeting)
}
