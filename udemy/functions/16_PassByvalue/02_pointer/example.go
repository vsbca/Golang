//this is using go programming language, integer

package main

import "fmt"

func main() {
	age := 44
	fmt.Println("This is the value of 'age' in main", age)
	changeme(&age) //This passes the address
	//changeme(age)  // pass by value 44
}

func changeme(z *int) {
	fmt.Println(z)
	fmt.Println(*z)
	*z = 24
	fmt.Println(z)
	fmt.Println(*z)
}
