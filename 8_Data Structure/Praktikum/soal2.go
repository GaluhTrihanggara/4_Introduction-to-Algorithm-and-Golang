package main

import "fmt"

func PairSum(array []int, target int) []int {
	n := len(array)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if array[i]+array[j] == target {
				return []int{i, j}
			}
		}
	}
	// jika tidak sesuai dengan target yang diinginkan maka
	return []int{0, 0}
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]
}
