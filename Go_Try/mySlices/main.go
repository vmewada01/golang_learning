package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList=	 []string{"apple","tomato","peach"}

	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fmt.Println(fruitList)

	fruitList= append(fruitList, "mango", "papaya")

	fmt.Println(fruitList)

	 fruitList = append(fruitList[1:3] )

	 fmt.Println(fruitList)

	 fruitList= append(fruitList[:3])
	 fmt.Println(fruitList)

     highScore := make([]int, 4)

	 highScore[0]= 92
	 highScore[1]= 45
	 highScore[2]= 67
	 highScore[3]= 89

     highScore = append(highScore, 111,222,333)

	 fmt.Println(sort.IntsAreSorted(highScore))

	 sort.Ints(highScore)


	  fmt.Println(highScore)


   var course = []string {"react js", "python", "swift", "ruby", "java","c++" , "c"}


   fmt.Println(course)

   var index int = 2

   course = append(course[:index], course[index+1:]...)

   fmt.Println(course)

  





	}