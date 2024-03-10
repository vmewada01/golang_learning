package main

import "fmt"

func main() {
	fmt.Println("welcome to loops in golang")

	// days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	// fmt.Println(days)
	// for i := 0; i < len(days); i++ {
    //     fmt.Println(days[i])
    // }

	// for j := range days{
	// 	fmt.Println(days[j])
	// }

	// for index, day := range days{
	// 	fmt.Printf("index is %v and value is %v\n", index,day)
	// }


	rougueValue := 1 

	for rougueValue <= 10 {
        
		if rougueValue==3 {
			goto lco
		}

		if rougueValue ==5 {
			rougueValue++
			continue
		}else if  rougueValue ==7 {
			break;
		}

        fmt.Printf("rougueValue is %v\n", rougueValue)
        rougueValue++
    }


	lco: 
	fmt.Println("Jumping at Learning golang")
}