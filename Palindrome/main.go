// Palindrome project main.go
package main

import (
	"fmt"
)

func palindrome(checkingNumber int) bool {
	var remainder, temp int
	var reverse int = 0

	temp = checkingNumber

	// For Loop used in format of While Loop
	for checkingNumber != 0 {
		remainder = checkingNumber % 10
		reverse = reverse*10 + remainder
		checkingNumber /= 10
	}

	if temp == reverse {
		return true
		// fmt.Printf("%d is a Palindrome",temp)
	} else {
		return false
		// fmt.Printf("%d is not a Palindrome",temp)
	}
}

func main() {
	var checkingNumber int
	fmt.Print("Enter any positive integer : ")
	fmt.Scan(&checkingNumber)

	if palindrome(checkingNumber) {
		fmt.Printf("%d is a Palindrome\n", checkingNumber)
	} else {
		fmt.Printf("%d is not a Palindrome\n", checkingNumber)
	}
}
