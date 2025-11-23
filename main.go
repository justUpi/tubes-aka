package main

import (
	"fmt"
	"time"
)

// TransposeIterative melakukan transpose matriks secara iteratif
func TransposeIterative(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return nil
	}

	transposed := make([][]int, len(matrix[0]))
	for i := range transposed {
		transposed[i] = make([]int, len(matrix))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}

// TransposeRecursive melakukan transpose matriks secara rekursif
func TransposeRecursive(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return nil
	}

	transposed := make([][]int, len(matrix[0]))
	for i := range transposed {
		transposed[i] = make([]int, len(matrix))
	}

	var transposeHelper func(int, int)
	transposeHelper = func(i, j int) {
		if i >= len(matrix) {
			return
		}
		if j >= len(matrix[i]) {
			transposeHelper(i+1, 0)
			return
		}
		transposed[j][i] = matrix[i][j]
		transposeHelper(i, j+1)
	}

	transposeHelper(0, 0)
	return transposed
}

func main() {
	const size = 1000
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
		for j := range matrix[i] {
			matrix[i][j] = i*size + j + 1
		}
	}

	const iterations = 100
	var totalDuration time.Duration

	// Measure iterative transpose
	for i := 0; i < iterations; i++ {
		start := time.Now()
		TransposeIterative(matrix)
		totalDuration += time.Since(start)
	}


	// Displaying partial result for iterative transpose
	transposedIterative := TransposeIterative(matrix)
	fmt.Println("\nHasil Transpose (Iteratif) - 5x5 pertama:")
	for i := 0; i < 5; i++ {
		fmt.Println(transposedIterative[i][:5])
	}
	averageTimeIterative := totalDuration / iterations
	fmt.Printf("Waktu Eksekusi Rata-rata (Iteratif): %v\n", averageTimeIterative)

	// Reset total duration for recursive transpose
	totalDuration = 0

	// Measure recursive transpose
	for i := 0; i < iterations; i++ {
		start := time.Now()
		TransposeRecursive(matrix)
		totalDuration += time.Since(start)
	}


	// Displaying partial result for recursive transpose
	transposedRecursive := TransposeRecursive(matrix)
	fmt.Println("\nHasil Transpose (Rekursif) - 5x5 pertama:")
	for i := 0; i < 5; i++ {
		fmt.Println(transposedRecursive[i][:5])
	}
	averageTimeRecursive := totalDuration / iterations
	fmt.Printf("Waktu Eksekusi Rata-rata (Rekursif): %v\n", averageTimeRecursive)
}