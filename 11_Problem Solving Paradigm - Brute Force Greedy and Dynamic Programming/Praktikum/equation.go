// program yang memecahkan x, y dan z untuk nilai yang diberikan A, B dan C. (1 ≤ A, B, C ≤ 10000).
// by Galuh Trihanggara
package main

import "fmt"

func SimpleEquations(a, b, c int) {
	var x, y, z int
	menemukan := false

	for x = -10000; x <= 10000; x++ {
		for y = -10000; y <= 10000; y++ {
			if x == y {
				continue
			}
			z = a - x - y
			if y == z || x == z {
				continue
			}
			if x*y*z == b && x*x+y*y+z*z == c {
				menemukan = true
				break
			}
		}
		if menemukan {
			break
		}
	}

	if menemukan {
		fmt.Printf("%d %d %d\n", x, y, z)
	} else {
		fmt.Println("no solution")
	}
}

func main() {
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
}
