package main

import "fmt"

func main() {
	fmt.Println(greet("Virendra", " Bisht"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
