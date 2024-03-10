package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("go lang file")

	content := "this needs to go in a file"

	file, err:=os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err)
	}
    length, err:=	  io.WriteString(file,content)
	checkNilErr(err)
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println("length is :",length)
 
	readFile("./mylcogofile.txt")

 defer	 file.Close()
}


func readFile(filename string){
   databyte, err :=  ioutil.ReadFile(filename)
  checkNilErr(err)
    // if err!= nil {
    //     panic(err)
    // }
    fmt.Println("Text data inside the file \n", string(databyte))

}

func checkNilErr(err error){
	if err!= nil {
        panic(err)
    }

}