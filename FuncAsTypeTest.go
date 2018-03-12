package main

import "fmt"

// Define function type
type testInt func(int) bool

func isOdd(interger int) bool {
	if interger%2 == 0 {
		return false
	}
	return true
}

func isEven(interger int) bool {
	if interger%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	// Do not define the length, It's the difference between array and slice
	slice := []int {1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("Slice = ", slice)
	fmt.Printf("\n")
	Odd := filter(slice, isOdd)
	fmt.Printf("Odd = ", Odd)
	fmt.Printf("\n")
	Even := filter(slice, isEven)
	fmt.Printf("Even = ", Even)
	fmt.Printf("\n")
}
