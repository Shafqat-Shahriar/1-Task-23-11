package main

import (
	"fmt"
	"math"
)

func main() {
	
	var length float64
	var width float64
	fmt.Print("Enter length of rectangle: ")
	fmt.Scanln(&length)
	fmt.Print("Enter width of rectangle: ")
	fmt.Scanln(&width)

	
	var radius float64
	fmt.Print("Enter radius of circle: ")
	fmt.Scanln(&radius)

	
	rectangleArea := length * width
	rectanglePerimeter := 2 * (length + width)

	
	circleArea := math.Pi * radius * radius
	circlePerimeter := 2 * math.Pi * radius

	
	fmt.Printf("Rectangle area: %.2f\n", rectangleArea)
	fmt.Printf("Rectangle perimeter: %.2f\n", rectanglePerimeter)
	fmt.Printf("Circle area: %.2f\n", circleArea)
	fmt.Printf("Circle circumference: %.2f\n", circlePerimeter)
}
