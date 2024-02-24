package main

import "fmt"

func main() {

	var x float64 // var x,y float64
	var y float64
	fmt.Print("Enter the numbers: \n X = ")
	fmt.Scanf("%f", &x)  // fmt.Scanf("%f %f", &x, &y)
	fmt.Print("\n Y = ")
	fmt.Scanf("%f", &y) 
	sum := x + y
	diff := x - y
	mul := x * y
	quot := x / y
	fmt.Println(sum, diff, mul, quot)
}
/*

Other commonly used format specifiers with fmt.Scanf include:

    %d: Scan an integer.
    %s: Scan a string.
    %c: Scan a single character.
    %t: Scan a boolean (true or false).
    %v: Scan the next value as a default representation.

*/
