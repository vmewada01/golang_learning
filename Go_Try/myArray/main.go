package main

import "fmt"

func main() {
	fmt.Println("hello world")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Banana"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegList= [5]string{"potato", "brokly", "tomato"}

 fmt.Println("values of veg List",vegList)
 fmt.Println(len(vegList))

}