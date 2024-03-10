package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://dummyjson.com/carts"

func main() {
	fmt.Println("web request")


	response, err := http.Get(url)
	
    if err!= nil {
        panic(err)
	}

	fmt.Println("response is of type",response)

	defer response.Body.Close()  //caller responsibility to close the tag

   databytes, err:= ioutil.ReadAll(response.Body)

   if err != nil {
	panic(err)
   }
   content := string(databytes)

   fmt.Println(content)
   
}