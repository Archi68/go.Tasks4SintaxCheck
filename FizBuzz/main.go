// FizBuzz project main.go
// The program should print the numbers from 1 to 100, except if the number is
// divisible by 3 then print ‘fizz’, if the number is divisible by 5 print ‘buzz’
// or if the number if divisible by both print ‘fizzbuzz’.

package main

import (
	"fmt"
)

func fizbuzz(upToNumber int) {
	for i := 1; i <= upToNumber; i++ {

		if i%3 == 0 {
			// Multiple of 3
			fmt.Printf("fizz")
		}
		if i%5 == 0 {
			// Multiple of 5
			fmt.Printf("buzz")
		}

		if i%3 != 0 && i%5 != 0 {
			// Neither, so print the number itself
			fmt.Printf("%d", i)
		}

		// A trailing new line (so both fizz + buzz can be printed on the same line)
		fmt.Printf(" ")
	}
	fmt.Println()
}
func main() {
	fizbuzz(100)
}
