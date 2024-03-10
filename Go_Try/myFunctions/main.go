package main

import "fmt"

func main() {
	fmt.Println("welcome to golang function")
	reader()
	readerTwo()

	result := addition(2,37)

	fmt.Println("Result of addition function",result)

	proResult  :=proAddition(1,2,3,4,5,6,4)

	fmt.Println("Result of proAddition function",proResult)
}

func proAddition(values ...int)int{
	total := 0
	for _, value := range values {
        total += value
    }
	return total
}


func addition(valOne int, valTwo int) int{
	return valOne + valTwo
}

func readerTwo(){
	fmt.Println("Another method")
}
func reader(){
	fmt.Println("Namaste from go lang")
}