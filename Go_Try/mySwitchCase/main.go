package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
    fmt.Println("Switch case in golang")

    rand.Seed(time.Now().UnixNano())

    diceNumber := rand.Intn(6) + 1

    fmt.Println("value of dice is ", diceNumber)

    switch diceNumber {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
        fallthrough // Move fallthrough here
    case 3:
        fmt.Println("three")
        fallthrough // Move fallthrough here
    case 4:
        fmt.Println("four")
    case 5:
        fmt.Println("five")
    case 6:
        fmt.Println("six")
    default:
        fmt.Println("default")
    }
}