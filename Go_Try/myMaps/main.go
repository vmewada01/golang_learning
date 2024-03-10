package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PHP"] = "PHP"
	languages["Java"] = "Java"
	languages["Python"] = "Python"

	fmt.Println("List of all langues",languages)
	fmt.Println("Js shorts for :", languages["JS"])

	delete(languages, "Java")
	fmt.Println("List of all languages",languages)

    for _, value := range languages{
		fmt.Printf("For key v, value is %v\n",  value)
	}

}