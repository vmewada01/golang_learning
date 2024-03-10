package main

import "fmt"



func main() {

	fmt.Println("Structs in golang")

	userDetails := User{"Vishal","vishal@gmail.com", true,21}

	fmt.Println(userDetails)

     fmt.Printf("Vishal details are: %+v\n",userDetails) 
	 fmt.Printf("Name is %v and email is %v.",userDetails.Name, userDetails.Email) 

}

type User struct {
	Name string
	Email string
	Status bool
	Age  int
}