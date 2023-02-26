// program untuk menghitung selisih absolut antara jumlah diagonalnya
// by Galuh Trihanggara
package main

import (
	"fmt"
	"math"
)

func perbedaanDiagonal(array [][]int) int {
	n := len(array)
	sumLeft, sumRight := 0, 0
	for i := 0; i < n; i++ {
		sumLeft += array[i][i]
		sumRight += array[i][n-i-1]
	}
	return int(math.Abs(float64(sumLeft - sumRight)))
}

func main() {
	array := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// Jika dihitung maka akan menghasilkan nilai selisih absolut antara jumlah diagonalnya yaitu 0
	// karena jika dihitung (1+5+9) - (3+5+7) = |15 - 15| = 0
	fmt.Println(perbedaanDiagonal(array))

	array = [][]int{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}
	// Jika dihitung maka akan menghasilkan nilai selisih absolut antara jumlah diagonalnya yaitu 2
	// karena jika dihitung (1+5+9) - (3+5+9) = |15 - 17| = 2
	fmt.Println(perbedaanDiagonal(array))
}
