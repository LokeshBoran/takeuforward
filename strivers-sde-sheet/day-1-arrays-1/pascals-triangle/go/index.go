package main

import "fmt"

func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	triangle[0] = []int{1}

	for i := 1; i < numRows; i++ {
		newElement := []int{1}
		previousRow := triangle[i-1]

		// Fill the rest of the row with sum of the squares above
		for j := 1; j < i; j++ {
			newElement = append(newElement, previousRow[j-1]+previousRow[j])
		}
		newElement = append(newElement, 1)
		triangle[i] = newElement
	}

	return triangle
}

func main() {

	fmt.Println("numRows=5", generate(5))

	fmt.Println("numRows=1", generate(1))
	fmt.Println("numRows=15", generate(15))
}
