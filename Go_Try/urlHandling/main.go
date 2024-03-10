package main

import (
	"fmt"
	"net/url"
)

const myUrl string= "https://dummyjson.com/carts?title=vishal"

func main() {
	fmt.Println("url handling")
	fmt.Println(myUrl)

	result , err := url.Parse(myUrl)

	if err!= nil {
        panic(err)
    }
	
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Println("The type of query params are %T\n",qparams)
	fmt.Println(qparams["title"])

	for _,val := range qparams{
		fmt.Println("param is :", val)
	}

	partsofUrl  := &url.URL{
		Scheme: "http",
		Host: "localhost",
		Path: "/carts",
		RawPath: "user=vishal",
        // RawQuery: url.Values{
        //     "title": []string{"vishal"},
        // }.Encode(),
	}

	anotherUrl := partsofUrl.String()

	fmt.Println(anotherUrl)
}