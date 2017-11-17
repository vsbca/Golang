package main

import (
	"fmt"
)

func main(){
	greet("virendra","Bisht")
}

func greet(fname, lname string){
	fmt.Println("Hello ",fname ,lname)
}
