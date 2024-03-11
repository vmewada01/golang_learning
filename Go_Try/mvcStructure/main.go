package main

import (
	"fmt"
	"mvcStructure/initializers"
)

func init(){

	initializers.ConnectToDatabase()
}

func main() {
	fmt.Println("hello world")
}