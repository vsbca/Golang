package main

import "fmt"

func hello() {
	fmt.Println("hello")
}

func world() {
	fmt.Println("Testing!!!!")
}
func main() {
	defer world()
	hello()
}
