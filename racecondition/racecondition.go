// increment() and print() share the same variable counter
// The outcome is non-deterministic, meaning print() may print outdated or incorrect values of counter because the order of access between the two goroutines is not controlled

package main

import (
	"fmt"
	"time"
)

var counter int

func increment() {
	for i := 0; i < 20; i++ {
		counter++
		time.Sleep(10 * time.Millisecond)
	}
}

func print() {
	for i := 0; i < 20; i++ {
		fmt.Println(counter)
		time.Sleep(10 * time.Millisecond)
	}
}
func main() {
	go increment()
	go print()

	time.Sleep(1 * time.Second)

	fmt.Println("Final counter value:", counter)
}
