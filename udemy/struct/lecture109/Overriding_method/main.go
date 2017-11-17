package main

import "fmt"

type person struct {
	Name string
	Age  int
}

//Promotions

type DoubleZero struct {
	person
	LicenseToKill bool
}

func (p person) Greetings() {
	fmt.Println("I'm just a regular greeting")
}

func (d DoubleZero) Greetings() {
	fmt.Println("I am greeting from DoubleZero")
}

func main() {
	p1 := DoubleZero{
		person: person{
			Name: "Virendra",
			Age:  20,
		},
		LicenseToKill: true,
	}
	p2 := person{"Ashima", 30}
	p2.Greetings()
	p1.Greetings()
}
