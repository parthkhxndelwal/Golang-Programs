package main

import "fmt"

func main() {
	slice1 := []int{10, 20, 30, 40, 50}
	slice2 := []int{5, 15, 25, 35, 45}

	if len(slice1) != len(slice2) {
		fmt.Println("Error: Slices must have the same length")
		return
	}

	result := make([]int, len(slice1))

	for i := 0; i < len(slice1); i++ {
		result[i] = slice1[i] + slice2[i]
	}

	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)
	fmt.Println("Result: ", result)
}