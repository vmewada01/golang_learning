package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("server get request")
	// PerformGetRequest()
	 PerformPostRequest()
	// PerformPostFormRequest()
}


func PerformPostRequest(){
	const myurl = "http://localhost:5000/users"

	requestBody := strings.NewReader(`
	
	{
        "firstName":"vishal",
		"lastName":"kumar",
        "phoneNumber":"0349857",
        "address":"sehore"
       
    }`)
    
    resp, err := http.Post(myurl,"application/json", requestBody)
    
    if err!= nil {
        panic(err)
    }
    defer resp.Body.Close()
    
    body, err := ioutil.ReadAll(resp.Body)
    
    if err!= nil {
        panic(err)
    }
    
    fmt.Println(string(body))
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

func  PerformPostFormRequest(){
	const myurl = "http://localhost:5000/users"

	data := url.Values{}

	data.Add("firstName","vishal")
	data.Add("lastName","kumar")
	data.Add("phoneNumber","0349857")
	data.Add("address","sehore")

        
    resp, err := http.PostForm(myurl, data)
    
    if err!= nil {
        panic(err)
    }
    defer resp.Body.Close()
    
    body, err := ioutil.ReadAll(resp.Body)
    
    if err!= nil {
        panic(err)
    }
    
    fmt.Println(string(body))	
}