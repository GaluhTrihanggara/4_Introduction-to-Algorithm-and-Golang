// program menemukan total maksimum jumlah bilangan dari deret sebuah integer secara berurutan.
// by Galuh Trihanggara
package main

import "fmt"

func MaxSequence(arr []int) int {
	maxjauh := 0
	maxberakhir := 0

	for _, num := range arr {
		maxberakhir += num

		if maxberakhir < 0 {
			maxberakhir = 0
		}

		if maxjauh < maxberakhir {
			maxjauh = maxberakhir
		}
	}

	return maxjauh
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))   // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))   // 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))    // 12
}
