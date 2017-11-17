package main

import "fmt"

//printing the array

func main() {
	//Array of type integers x
	var x [256]int

	fmt.Println(len(x))
	//storing the value in array
	for i := 0; i < 256; i++ {
		x[i] = i
	}
	//print the values
	for i, v := range x {
		fmt.Println("%v - %T - %b \n", v, v, v)
		if i > 50
		   break
	}

}
