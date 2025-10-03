package main

import (
	"fmt"
	"os"
	"strconv"
)

func factorial(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 || n == 1 {
		return 1
	}

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <number>")
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: Invalid number format")
		return
	}

	if num < 0 {
		fmt.Println("Error: Factorial is not defined for negative numbers")
		return
	}

	result := factorial(num)
	fmt.Printf("Factorial of %d is %d\n", num, result)
}