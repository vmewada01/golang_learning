package main

import "fmt"

func main() {
	defer fmt.Println("welcome to differ part")

	defer example()
	fmt.Println("Exploring Defer function")

	myDefer()

}

func example(){
	fmt.Println("example for this differ")
}
func myDefer(){
	for i:=0; i<5; i++{
		defer fmt.Println(i)
	}
}