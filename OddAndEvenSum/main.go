// OddAndEvenSum project main.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var sumEven, sumOdd int = 0, 0
	for i := 0; i < 10; i += 1 {
		rand.Seed(time.Now().UnixNano())
		number := rand.Intn(51) + 10
		fmt.Print(number)
		fmt.Print(" ")
		if number%2 == 0 {
			// Even
			sumEven += number
		} else {
			sumOdd += number
		}
	}
	fmt.Print("\n")
	fmt.Println("Even numbers sum are: ", sumEven)
	fmt.Println("Odd numbers sum are: ", sumOdd)
}
