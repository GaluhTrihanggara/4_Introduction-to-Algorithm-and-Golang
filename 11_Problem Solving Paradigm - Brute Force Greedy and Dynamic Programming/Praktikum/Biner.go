// Program Represntasi Biner
// by Galuh Trihanggara
package main

import (
	"fmt"
	"strconv"
)

// pertama membuat representasi binary terlebih dahulu
func getRepresntasiBinary(n int) string {
	binary := strconv.FormatInt(int64(n), 2)
	for len(binary) < 1 {
		binary = "0" + binary
	}
	return binary
}

// kemudian membuat binary array
func getBinaryArray(n int) []string {
	result := make([]string, n+1)
	for i := 0; i <= n; i++ {
		result[i] = getRepresntasiBinary(i)
	}
	return result
}

// mencetak hasilnya
func main() {
	fmt.Println(getBinaryArray(2)) // [0 1 10]
	fmt.Println(getBinaryArray(3)) // [0 1 10 11 100 101]
}
