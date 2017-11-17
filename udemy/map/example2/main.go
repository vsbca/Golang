package main

import "fmt"

func main() {

	//declaration of map with composite literal
	//MAP is dictionary with Key value pair
	myGreeting := map[string]string{
		"Tim":   "Good morning",
		"Jenny": "Bonjour",
	}

	fmt.Println(myGreeting["Jenny"])

}
