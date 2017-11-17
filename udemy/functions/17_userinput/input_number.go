package main

import "fmt"

func checktheinput() {
	var smallnumber int
	var bignumber int
	fmt.Println("Please enter the number!!")
	fmt.Scan(&smallnumber)
	fmt.Scan(&bignumber)

	fmt.Println("These are the input smallnumber and  bignumber", smallnumber, "and", bignumber)
}

func main() {
	checktheinput()
}
