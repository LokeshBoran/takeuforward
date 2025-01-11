package main

import "fmt"

func setZeroes(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])

	hasFirstColZero := false

	// Step 1: Traverse the matrix and mark 1st row & col accordingly:
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				// Mark i-th row:
				matrix[i][0] = 0

				// Mark j-th column:
				if j != 0 {
					matrix[0][j] = 0
				} else {
					hasFirstColZero = true
				}
			}
		}
	}

	// Step 2: Mark with 0 from (1,1) to (n-1, m-1):
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] != 0 {
				// Check for col & row:
				if matrix[i][0] == 0 || matrix[0][j] == 0 {
					matrix[i][j] = 0
				}
			}
		}
	}

	// Step 3: Finally mark the 1st col & then 1st row:
	if matrix[0][0] == 0 {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}
	if hasFirstColZero {
		for i := 0; i < n; i++ {
			matrix[i][0] = 0
		}
	}

	fmt.Println(fmt.Sprintf("%v", matrix))

}

func main() {
	array := [][]int{}

	array = [][]int{{0, 1}}
	fmt.Println("array=>", fmt.Sprintf("%v", array))
	setZeroes(array)
	array = [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	fmt.Println("array=>", fmt.Sprintf("%v", array))
	setZeroes(array)
	array = [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	fmt.Println("array=>", fmt.Sprintf("%v", array))
	setZeroes(array)
}
