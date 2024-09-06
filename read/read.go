package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define the Name struct with two fields
type Name struct {
	fname string
	lname string
}

func main() {
	fmt.Print("Enter the name of the text file: ")
	var filename string
	_, err := fmt.Scanln(&filename)
	if err != nil {
		fmt.Println("Error reading filename:", err)
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var names []Name

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid line format, skipping:", line)
			continue
		}

		fname := parts[0]
		lname := parts[1]

		names = append(names, Name{fname: fname, lname: lname})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("\nNames found in the file:")
	for _, name := range names {
		fmt.Printf("First Name: %s, Last Name: %s\n", name.fname, name.lname)
	}
}
