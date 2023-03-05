// Program ini menerima money sebagai parameter pertama dan yang kedua berupa list harga produk yang dapat dibeli.
// by Galuh Trihanggara
package main

import (
	"fmt"
	"sort"
)

func MaximumBuyProduct(money int, productPrice []int) {
	// mengurutkan harga produk dalam urutan menaik
	sort.Ints(productPrice)

	// Melakukan iterate melalui harga produk yang diurutkan
	numOfProducts := 0
	for _, price := range productPrice {
		if price <= money {
			money -= price
			numOfProducts++
		} else {
			break
		}
	}

	fmt.Println(numOfProducts)
}

func main() {
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})      // 3
	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000}) // 4
	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})   // 4
	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})           // 1
	MaximumBuyProduct(0, []int{10000, 30000})                        // 0
}
