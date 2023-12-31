package main

import "fmt"

func main() {

    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "apple"
    fmt.Println("Initial value of f:", f)

    // Change the value of f
    f = "banana"
    fmt.Println("Updated value of f:", f)
    
    /* 
    Using " := " is a concise way of declaring and initializing variables, 
    especially when the type can be inferred from the right-hand side 
    of the assignment. It's commonly used in shorter variable 
    declarations within functions or local scopes. */
}
