package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 0, 3)

	for {
		var input string
		fmt.Printf("Please enter a number (or 'X' to exit): ")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error. Please try again.")
			continue
		}
		if input == "X" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid number.")
			continue
		}

		sli = append(sli, num)
		sort.Ints(sli)
		fmt.Println("Sorted slice: ", sli)

		if len(sli) == cap(sli) {
			break
		}
	}
}
