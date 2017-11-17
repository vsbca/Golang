package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

//Struct aggregate type

func main() {
	p1 := person{"Virendra", "Singh", 20}
	fmt.Println(p1.first, p1.age)
}
