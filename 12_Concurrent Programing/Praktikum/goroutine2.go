// program yang mencetak bilangan kelipatan 3 dengan menerapkan goroutine dan channel!
// by Galuh Trihanggara
package main

import (
	"fmt"
	"time"
)

func main() {
	// membuat channel untuk mengirim dan menerima bilangan kelipatan 3
	ch := make(chan int)

	// menjalankan goroutine untuk mencetak bilangan kelipatan 3 ke channel
	go func() {
		for i := 1; ; i++ {
			if i%3 == 0 {
				ch <- i
			}
		}
	}()

	// mencetak bilangan kelipatan 3 dari channel
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
		time.Sleep(3 * time.Second)
	}
}
