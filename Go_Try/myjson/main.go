package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"coursename"`
	Price       int
	Platform  string  `json:"website"`
	Password string `json:"-"`
	Tags []string  `json:"tags,omitempty"`
}

func main() {
	fmt.Println("go lang json")
	EncodeJson()
}

func EncodeJson() {
	userDetails := []course{
		{
			Name:     "ReactJs",
			Price:    223,
			Platform: "google.com",
			Password: "abcd123",
			Tags:     []string{"hello", "23"},
		},
		{
			Name:     "Mern",
			Price:    223,
			Platform: "google.com",
			Password: "abcd123",
			Tags:     []string{"world", "12"},
		},
		{
			Name:     "angular",
			Price:    223,
			Platform: "google.com",
			Password: "abcd123",
			Tags:     []string{"new", "2"},
		},
	}

	finalJson, err := json.Marshal(userDetails)
	
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("%s\n", finalJson)
	final_indent_json , err := json.MarshalIndent(userDetails,"","\t")

	if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", final_indent_json)
}