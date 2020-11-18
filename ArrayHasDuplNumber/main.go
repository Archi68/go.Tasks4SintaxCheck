// ArrayHasDuplNumber project main.go
package main

import (
	"fmt"
)

func main() {
	var numbers = [...]int{1, 2, 3, 4, 5, 12, 32}
	var hasDuplicate bool = false
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] == numbers[j] {
				hasDuplicate = true
			}
		}
	}
	if hasDuplicate {
		fmt.Println("Array has duplicate")
	} else {
		fmt.Println("There is no duplicates in array")
	}
}
