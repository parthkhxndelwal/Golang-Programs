package main

import "fmt"

func main() {
	const size = 5

	array1 := [size]int{10, 20, 30, 40, 50}
	array2 := [size]int{5, 15, 25, 35, 45}
	var result [size]int

	for i := 0; i < size; i++ {
		result[i] = array1[i] + array2[i]
	}

	fmt.Println("Array 1:", array1)
	fmt.Println("Array 2:", array2)
	fmt.Println("Result: ", result)
}