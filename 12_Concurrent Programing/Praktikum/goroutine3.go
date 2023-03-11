// Program yang mencetak bilangan kelipatan 3 dengan menerapkan goroutine dan buffer channel!
// by Galuh Trihanggara
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 10)

	go func() {
		for i := 1; i <= 20; i++ {
			if i%3 == 0 {
				ch <- i
			}
		}
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}
}
