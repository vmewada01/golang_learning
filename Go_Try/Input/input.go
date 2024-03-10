package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
   welcome := "welcome to go lang"

    fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	name, _ := reader.ReadString('\n')

	fmt.Println("Thank for name , ", name)
	fmt.Printf("Type of this name is  %T",name)

}