//calculate factorial of a number

package main

import "fmt"

func fact(x int) int {
	if x == 0 {
		return 1
	} else {
		return (x * fact(x-1))
	}
}

func main() {

	fmt.Println(fact(2))
}
