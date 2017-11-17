package main

import "fmt"

func main() {

	//Definition of map
	myGreeting := map[int]string{
		0: "Good morning",
		1: "Bonjour!!!",
		2: "Buenos dias",
		3: "Shubh Prabhat!!!",
	}

	for key, val := range myGreeting {
		fmt.Println(key, "is", val)
	}
}
