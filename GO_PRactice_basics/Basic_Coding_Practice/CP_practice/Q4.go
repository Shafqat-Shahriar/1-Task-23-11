package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	if number % 2 == 0 {
		fmt.Println(number, "is even.")
	} else {
		fmt.Println(number, "is odd.")
	}
}
