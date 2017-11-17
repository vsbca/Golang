//Go static type

package main

import "fmt"

type foo int

func main() {

	var myage foo

	myage = 44
	fmt.Println("This is myAge value:", myage)

	var yourAge int
	yourAge = 29

	fmt.Println("This is yourAge:", yourAge)
	//This should not work
	fmt.Println(myage + yourAge)

}
