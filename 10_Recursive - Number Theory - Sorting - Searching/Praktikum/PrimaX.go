// Program Mencari Bilangan Prima
// by Galuh Trihanggara
package main

import "fmt"

func primeX(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func getPrime(n int) int {
	count, hasil := 0, 0
	for count < n {
		hasil++
		if primeX(hasil) {
			count++
		}
	}
	return hasil
}

func main() {
	fmt.Println(getPrime(1))  // 2
	fmt.Println(getPrime(5))  // 11
	fmt.Println(getPrime(8))  // 19
	fmt.Println(getPrime(9))  // 23
	fmt.Println(getPrime(10)) // 29
}
