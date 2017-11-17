package main

import "fmt"

func main() {
	z := "Virendra"
	changeme(&z)
	fmt.Println("This is the value of z in main", z)
}

func changeme(x *string) {
	*x = "Ashima"
	fmt.Println(*x)
}
