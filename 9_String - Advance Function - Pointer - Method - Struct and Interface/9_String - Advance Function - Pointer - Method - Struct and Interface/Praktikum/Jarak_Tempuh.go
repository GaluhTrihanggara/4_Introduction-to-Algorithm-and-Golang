// Program Menghitung seberapa jauh yang akan ditempuh jika bahan bakar terisi 1500ml(1.5l)/Mill
// by Galuh Trihanggara
package main

import (
	"fmt"
)

type Car struct {
	typeCar string
	fuelIn  float64 // Bahan bakar dalam mililiter
}

// Method getJarakTempuh() memiliki receiver c yang merupakan objek dari struct Car.
func (c Car) getJarakTempuh() float64 {
	// Konversi bahan bakar dari ml ke liter
	fuelInL := c.fuelIn / 1000

	// menghitung perkiraan jarak yang bisa ditempuh berdasarkan bahan bakar yang terisi
	jarakTempuh := fuelInL / 1.5

	return jarakTempuh
}

func main() {
	// Contoh penggunaan
	myCar := Car{typeCar: "BMW", fuelIn: 1500} // Bahan bakar 1500 mL (1.5 L)
	fmt.Println("Jadi jarak yang bisa ditempuh berdasarkan bahan bakar yang sedang terisi yaitu", myCar.getJarakTempuh(), "mill")
}
