package main

import "fmt"



func main() {

	fmt.Println("methods in golang")

	userDetails := User{"Vishal","vishal@gmail.com", true,21}

	fmt.Println(userDetails)

     fmt.Printf("Vishal details are: %+v\n",userDetails) 
	// fmt.Printf("Name is %v and email is %v\n",userDetails.Name, userDetails.Email) 

	  userDetails.GetStatus()
	  userDetails.NewMail()	
}

type User struct {
	Name string
	Email string
	Status bool
	Age  int
}

func (u User) GetStatus(){
	fmt.Println("is User active", u.Status)
}

func (u User) NewMail(){
	u.Email= "test@gmail.com"
	fmt.Println("Email of this user is", u.Email)
}