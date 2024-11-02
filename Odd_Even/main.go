package main

import "fmt"

func main() {
	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range intSlice {
		if number%2 != 0 {
			fmt.Println(number, " is odd")
		} else {
			fmt.Println(number, " is even")
		}
	}
}
