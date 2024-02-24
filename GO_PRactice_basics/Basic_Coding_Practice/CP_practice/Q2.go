package main

import (
	"fmt"
	"strings"
)


func isPalindrome(text string) bool{
	if len(text) <=1{
	    return true
	}

	// Converting to lowercase and removing non-alphanumeric characters

	text = strings.ToLower(strings.ReplaceAll(text, `[^a-zA-Z0-9]`, ""))

	
	// using two pointers to compare characters from begining and end

	left, right := 0, len(text)-1
	for left < right {
		if text[left] != text[right] {
			return false
		}
		left++
		right--
	}

	return true
}





func main() {
	var text string
	fmt.Println("Enter a string: ")
	fmt.Scanln(&text)

	if isPalindrome(text) {  
		fmt.Println(text, "is a palindrome!")
	} else {
		fmt.Println(text, "is not a palindrome.")
	}
}
