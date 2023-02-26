//program yang dapat mencari angka yang hanya muncul sekali pada input yang diberikan
// by Galuh Trihanggara

package main

import (
	"fmt"
)

func munculSekali(angka string) []string {
	mapDict := make(map[rune]int)

	//  cara menghitung frekuensi kemunculan setiap karakter/angka pada string
	for _, char := range angka {
		mapDict[char]++
	}
	var result []string

	for _, char := range angka {
		if mapDict[char] == 1 {
			result = append(result, string(char))
		}
	}
	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6, 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}
