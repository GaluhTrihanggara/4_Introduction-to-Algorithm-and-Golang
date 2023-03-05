// Program mengurutkan barang berdasarkan jumlah kemunculanya
// Galuh Trihanggara
package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(item []string) []pair {
	counter := make(map[string]int)
	for _, v := range item {
		counter[v]++
	}

	pairs := make([]pair, len(counter))
	i := 0
	for k, v := range counter {
		pairs[i] = pair{k, v}
		i++
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	return pairs
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	// [{js 4} {ruby 2} {golang 1}]
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// [{A 4} {B 3} {D 2} {C 1}]
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
	// [{basketball 1} {football 1} {tenis 1}]
}
