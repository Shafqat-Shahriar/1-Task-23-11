package main

import "fmt"

func main() {
	fmt.Println("Booleans \nExercise")

	fmt.Println(true && false) // true
	fmt.Println(true || false) // true
	fmt.Println(!true)         // false

	var b int = 7
	var h int = 1

	//var AT float64 = (0.5 * b * h)     !!!!  this wont work because

	/*
		    	 ++++++++++  Note  ++++++++++

		the division 1 / 2 would still be performed using integer arithmetic,
		and the fractional part would be truncated. Which can lead to inaccurate results,
		especially when dealing with calculations that involve fractions or decimals.

		In this code, AT will be calculated as 0 * 7 * 1, which is 0.
	*/

	var AT float64 = 0.5 * float64(b) * float64(h) //  correct way

	fmt.Println("\nArea of triangle =", AT) //integer and string
	fmt.Println("8.0/3.0 =", 8.0/3.0)

}
