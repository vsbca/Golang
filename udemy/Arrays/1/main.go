//This is the function to make an array

package main

import "fmt"

func main() {

	//decalared array a of 5 integers, initially all items will be 0
	var a [5]int
	fmt.Println("emp ", a)
	// set 4th element as 100
	a[4] = 100
	fmt.Println("set: ", a)

	//declare new array, with predefined values
	b := [5]int{1, 2, 3, 4, 5}

	fmt.Println("This is b:", b)

	//Declare two dimensional array
	var twoD [2][3]int
	//For loop doesn't need ()
	fmt.Println("Before :", twoD)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
