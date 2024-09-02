package main

import (
	"fmt"
)

func main() {
	var truncNum float32
	var intNum int

	fmt.Println("Please enter the float number you want to truncate:")
	fmt.Printf(">> ")
	_, err := fmt.Scan(&truncNum)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	intNum = int(truncNum)
	fmt.Printf("Truncated integer: %d\n", intNum)
}
