package main

import "fmt"

func main() {
var primeNumbers = []int{2,5,7}
primeNumbers = append(primeNumbers,11)

fmt.Println(primeNumbers)

var Doxc [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            Doxc[i][j] = i + j
        }
    }
    fmt.Println("2d: ", Doxc)

  
}


