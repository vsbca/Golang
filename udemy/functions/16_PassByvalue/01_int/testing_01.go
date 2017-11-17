package main

//This is the example of x not being changed using function call ( Pass by value)
import "fmt"

func changeme(x int) {
	fmt.Println("this is the value of x", x)
	x = 24
}

func main() {
	age := 40
	changeme(age)
	fmt.Println(age)
}
