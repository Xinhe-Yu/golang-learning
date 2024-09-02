package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Println("Please enter your input:")
	fmt.Printf(">> ")
	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") && strings.Contains(input, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
