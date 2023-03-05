// program untuk menentukan substring yang sama di antara keduanya, dengan diberikan dua string A dan B
// by Galuh Trihanggara
package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	if strings.HasPrefix(a, b) {
		return b
	} else if strings.HasPrefix(b, a) {
		return a
	} else {
		return ""
	}

}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}
