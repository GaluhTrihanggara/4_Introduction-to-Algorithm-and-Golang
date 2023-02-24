package main

import (
	"fmt"
	"math"
)

func primeNumber(n int) bool {
	if n <= 1 {
		return false
	}
	sq_root := int(math.Sqrt(float64(n)))
	for i := 2; i <= sq_root; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeNumber(1000000007))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}
