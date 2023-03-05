// Program mengenkripsi(encrypt) sebuah nama dan dapat dideskripsi(decrypt)
// Galuh Trihanggara
package main

import "fmt"

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""

	for _, char := range s.name {
		if char == 'z' {
			nameEncode += "a"
		} else {
			nameEncode += string(char + 1)
		}
	}

	s.nameEncode = nameEncode
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""

	for _, char := range s.nameEncode {
		if char == 'a' {
			nameDecode += "z"
		} else {
			nameDecode += string(char - 1)
		}
	}

	return nameDecode
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of student’s name " + a.nameEncode + " is : " + c.Decode())
	}
}
