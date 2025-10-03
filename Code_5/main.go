package main

import (
	"fmt"
	"os"
	"strconv"
)

func factorialRecursive(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorialRecursive(n-1)
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

	result := factorialRecursive(num)
	fmt.Printf("Factorial of %d is %d\n", num, result)
}