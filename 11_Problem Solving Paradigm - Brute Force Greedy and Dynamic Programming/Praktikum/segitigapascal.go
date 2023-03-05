// Program Segitiga Pascal
// by Galuh Trihanggara
package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for i := range triangle {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][i] = 1, 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	return triangle
}

func main() {
	numRows := 5
	pascalTriangle := generate(numRows)
	fmt.Println(pascalTriangle)
}
