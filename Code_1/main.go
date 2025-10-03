package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <num1> <operator> <num2>")
		fmt.Println("Operators: +, -, *, /")
		return
	}

	num1, err1 := strconv.ParseFloat(os.Args[1], 64)
	operator := os.Args[2]
	num2, err2 := strconv.ParseFloat(os.Args[3], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Error: Invalid number format")
		return
	}

	var result float64

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Error: Invalid operator. Use +, -, *, or /")
		return
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}