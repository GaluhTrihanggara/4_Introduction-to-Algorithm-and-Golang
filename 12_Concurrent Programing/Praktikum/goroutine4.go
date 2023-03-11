// program yang yang menerapkan mutex! Jenis program yang dibuat bebas ( Kasus perhitungan faktorial).
// by Galuh Trihanggara
package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var result int64 = 1

	done := make(chan bool)

	for i := 1; i <= 10; i++ {
		go func(n int) {
			mu.Lock()
			result *= int64(n)
			mu.Unlock()

			if n == 10 {
				done <- true
			}
		}(i)
	}

	<-done
	fmt.Printf("Hasil faktorial dari 10 adalah %d\n", result)
}
