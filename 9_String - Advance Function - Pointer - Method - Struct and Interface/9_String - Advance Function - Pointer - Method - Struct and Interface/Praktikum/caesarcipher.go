// program caesar cipher
// by Galuh Trihanggara
package main

import "fmt"

func caesar(offset int, input string) string {
	var result string
	for _, char := range input {
		if char == ' ' {
			result += " "
		} else {
			// konversi rune menjadi integer untuk mendapatkan nilai ASCII
			charInt := int(char)
			// lakukan pergeseran dengan mengurangi nilai ASCII huruf 'a' dan menambahkan offset
			shiftedInt := (charInt - 97 + offset) % 26
			// tambahkan nilai ASCII huruf 'a' dan konversi kembali ke rune
			shiftedChar := rune(shiftedInt + 97)
			result += string(shiftedChar)
		}
	}
	return result
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cnvc
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
