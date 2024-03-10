package main

import "fmt"

func main() {
	fmt.Println("hello world")

	// var ptr *int 

	// fmt.Println("Value of ptr is",ptr)


	myNumber := 23 

	var pointer = &myNumber

	fmt.Println("Value of myNumber is",pointer)
	fmt.Println("Value of myNumber is",*pointer)

	*pointer = *pointer + 2 

	fmt.Println("Value of myNumber is",myNumber)
}