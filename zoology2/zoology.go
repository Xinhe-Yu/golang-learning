package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

type Bird struct{}

type Snake struct{}

func (c Cow) Eat() { fmt.Println("grass") }

func (c Cow) Move() { fmt.Println("walk") }

func (c Cow) Speak() { fmt.Println("moo") }

func (b Bird) Eat() { fmt.Println("worms") }

func (b Bird) Move() { fmt.Println("fly") }

func (b Bird) Speak() { fmt.Println("peep") }

func (s Snake) Eat() { fmt.Println("mice") }

func (s Snake) Move() { fmt.Println("slither") }

func (s Snake) Speak() { fmt.Println("hsss") }

func main() {
	animals := make(map[string]Animal)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		scanner.Scan()
		input := scanner.Text()

		parts := strings.Split(input, " ")

		if len(parts) != 3 {
			fmt.Println("Invalid command format. Please enter 3 words.")
			continue
		}

		command := parts[0]
		name := parts[1]
		param := parts[2]

		switch command {
		case "newanimal":
			switch param {
			case "cow":
				animals[name] = Cow{}
			case "bird":
				animals[name] = Bird{}
			case "snake":
				animals[name] = Snake{}
			default:
				fmt.Println("Unknown animal type. Please use 'cow', 'bird', or 'snake'.")
				continue
			}
			fmt.Println("Created it!")
		case "query":
			animal, exists := animals[name]
			if !exists {
				fmt.Println("No animal with that name found.")
				continue
			}

			switch param {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Unknown query. Please use 'eat', 'move', or 'speak'.")
			}
		default:
			fmt.Println("Unknown command. Please use 'newanimal' or 'query'.")
		}
	}
}
