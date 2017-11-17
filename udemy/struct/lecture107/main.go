package main

//Go OOP
//Encapsulation
/*
state
behaviour
exported/un-exported

Reusabiity
Polymorphism
interface

*/

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	Person
	First         string
	LicenseToKill bool
}

//Composition with go Ardan

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}

	fmt.Println("This is the value", p1.First, p1.Person.First)
}
