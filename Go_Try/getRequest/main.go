package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("server get request")
	PerformGetRequest()
}

func PerformGetRequest(){
	const myurl = "http://localhost:5000/"
    
	resp, err := http.Get(myurl)

	if err!= nil {
        panic(err)
    }

	fmt.Println("status code :", resp.StatusCode)
    fmt.Println("content length :", resp.ContentLength)
	content, _ := ioutil.ReadAll(resp.Body)

	// fmt.Println(string(content))

	var responseString strings.Builder

	byteCount , _ := responseString.Write(content)

	fmt.Println("Byte count is",byteCount)
	fmt.Println(responseString.String())
	defer resp.Body.Close()
}