// program untuk mengkonversi dari angka normal ke Angka Romawi!
// by Galuh Trihanggara
package main

import "fmt"

func toRoman(num int) string {
	var roman string
	for num > 0 {
		switch {
		case num >= 1000:
			roman += "M"
			num -= 1000
		case num >= 900:
			roman += "CM"
			num -= 900
		case num >= 500:
			roman += "D"
			num -= 500
		case num >= 400:
			roman += "CD"
			num -= 400
		case num >= 100:
			roman += "C"
			num -= 100
		case num >= 90:
			roman += "XC"
			num -= 90
		case num >= 50:
			roman += "L"
			num -= 50
		case num >= 40:
			roman += "XL"
			num -= 40
		case num >= 10:
			roman += "X"
			num -= 10
		case num >= 9:
			roman += "IX"
			num -= 9
		case num >= 5:
			roman += "V"
			num -= 5
		case num >= 4:
			roman += "IV"
			num -= 4
		default:
			roman += "I"
			num--
		}
	}
	return roman
}

func main() {
	fmt.Println(toRoman(4))    // IV
	fmt.Println(toRoman(9))    // IX
	fmt.Println(toRoman(23))   // XXIII
	fmt.Println(toRoman(2021)) // MMXXI
	fmt.Println(toRoman(1646)) // MDCXLVI
}
