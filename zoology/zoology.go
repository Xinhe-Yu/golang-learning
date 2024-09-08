package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	animals := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter an animal and a request (e.g., 'cow eat').")

	for {
		fmt.Print("> ")

		scanner.Scan()
		input := scanner.Text()

		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("Invalid input. Please enter an animal and a request (e.g., 'cow eat').")
			continue
		}

		animalName := parts[0]
		request := parts[1]

		animal, exists := animals[animalName]
		if !exists {
			fmt.Println("Unknown animal. Available animals are: cow, bird, snake.")
			continue
		}

		switch request {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Invalid request. Available requests are: eat, move, speak.")
		}
	}
}
