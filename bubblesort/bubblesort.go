package main

import (
	"fmt"
)

func Swap(slice []int, i int) {
	slice[i], slice[i+1] = slice[i+1], slice[i]
}

func BubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}
}

func main() {
	const maxNumbers = 10
	numbers := make([]int, 0, maxNumbers)

	fmt.Printf("Enter up to %d integers. Type 'done' when finished:", maxNumbers)

	for len(numbers) < maxNumbers {
		var input string
		fmt.Printf("Enter integer %d: ", len(numbers)+1)
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		if input == "done" {
			break
		}

		var number int
		_, err = fmt.Sscanf(input, "%d", &number)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			continue
		}

		numbers = append(numbers, number)
	}

	BubbleSort(numbers)

	fmt.Print("Sorted integers: ")
	for i, num := range numbers {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(num)
	}
	fmt.Println()
}
