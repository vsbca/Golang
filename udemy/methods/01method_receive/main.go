package main

import (
	"fmt"
)

type User struct {
	FirstName, LastName string
}

//This is the example of method
//The receiver is defined int he beginning of the function

func (u *User) Greetings() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

func main() {
	u := &User{"Ashima", "Virendra"}
	fmt.Println(u.Greetings())
}
