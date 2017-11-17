package main

import "fmt"

func main() {
	fmt.Println(greet("Jane", "Doe"))
}

//func <function_name> variable return type
func greet(fname, lname string) (s string) {
	s = fmt.Sprintln(fname, lname)
	return
}
