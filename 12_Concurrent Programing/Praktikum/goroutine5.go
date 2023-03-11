// program untuk Hitung frekuensi huruf dalam teks menggunakan perhitungan paralel (Bersamaan).
// by Galuh Trihanggara
package main

import (
	"fmt"
	"sync"
)

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
	freq := make(map[rune]int)
	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(len(text))

	for _, char := range text {
		go func(c rune) {
			mu.Lock()
			freq[c]++
			mu.Unlock()
			wg.Done()
		}(char)
	}

	wg.Wait()

	for char, count := range freq {
		fmt.Printf("%c: %d\n", char, count)
	}
}
