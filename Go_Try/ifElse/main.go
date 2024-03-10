package main

import "fmt"

func main() {
	fmt.Println("If Else statement")

	loginCount := 10 
	 
	var result string

	if loginCount < 10 {
		result = "regular user"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "exact 10"
	}
     fmt.Println(result)


	 if 9%2 == 0{
        fmt.Println("it is even")
	 }else {
		fmt.Println("it is odd")
	 }

	 if  num:=31 ; num <10 {
		fmt.Println("less than 10")
	 }else{
		fmt.Println("more than 10")
	 }
}