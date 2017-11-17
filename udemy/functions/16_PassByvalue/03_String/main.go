package main

import "fmt"

func main() {
	x := "Virendra"
	changeme(x)
	fmt.Println("This is the value of 'x' in main", x)
}

func changeme(z string) {
	z = "Ashima"
	fmt.Println("This is 'z' in changeme", z)
}
