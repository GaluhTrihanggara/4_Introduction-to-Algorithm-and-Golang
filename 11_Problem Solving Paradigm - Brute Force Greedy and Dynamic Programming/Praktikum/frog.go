// Program menghitung biaya total minimum yang mungkin dikeluarkan sebelum katak mencapai batu N
// by Galuh Trihanggara
package main

import "fmt"

func Frog(jumps []int) int {
	jumlah := len(jumps)
	if jumlah == 0 {
		return 0
	}
	if jumlah == 1 {
		return jumps[0]
	}

	// inisialisasi nilai awal
	dp := make([]int, jumlah)
	dp[0] = 0
	dp[1] = abs(jumps[1] - jumps[0])

	// iterasi untuk mencari nilai minimum
	for i := 2; i < jumlah; i++ {
		dp[i] = min(dp[i-1]+abs(jumps[i]-jumps[i-1]), dp[i-2]+abs(jumps[i]-jumps[i-2]))
	}

	return dp[jumlah-1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
