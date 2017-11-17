package main

import "fmt"

type personal struct {
	name string
	last string
	age  int
}

//Declare a method
//Method has receiver

func (c personal) fullname() string {
	return (c.name + " " + c.last)
}

func main() {
	p1 := personal{"Virendra", "Singh", 21}
	fmt.Println(p1.fullname())
}
