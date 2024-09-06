package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	personInfo := make(map[string]string)

	var inputBuffer [100]byte

	fmt.Print("Enter your name: ")
	n, err := os.Stdin.Read(inputBuffer[:])
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}
	name := strings.TrimSpace(string(inputBuffer[:n]))
	personInfo["name"] = name

	fmt.Print("Enter your address: ")
	n, err = os.Stdin.Read(inputBuffer[:])
	if err != nil {
		fmt.Println("Error reading address:", err)
		return
	}
	address := strings.TrimSpace(string(inputBuffer[:n]))
	personInfo["address"] = address

	jsonData, err := json.Marshal(personInfo)
	if err != nil {
		fmt.Println("Error creating JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
