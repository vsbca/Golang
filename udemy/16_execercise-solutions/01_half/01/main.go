package main


import "fmt"

func half(n int) (int, bool){
	return n/2, n%2 == 0
}

func main(){
	h, event := half(5)
	fmt.
}