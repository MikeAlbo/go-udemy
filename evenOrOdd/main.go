package main

import (
	"fmt"
	"strconv"
)

// create a main function
// create a slice of ints from 0 to 10
// iterate thorugh a slice with a for loop
// 		if the value is even, print out "is even"
//		if the value is odd, print out "is odd"
// example "0 is even"

func main() {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, int := range ints {
		if int%2 == 0 {
			fmt.Println(strconv.Itoa(int) + " is even")
		} else {
			fmt.Println(strconv.Itoa(int) + " is odd")
		}
	}
}
