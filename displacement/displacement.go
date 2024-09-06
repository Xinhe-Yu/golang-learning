package main

import (
	"fmt"
)

func GenDisplaceFn(a, vo, so float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + vo*t + so
	}
}

func main() {
	var a, vo, so, t float64
	a = 10
	vo = 2
	so = 1

	fn := GenDisplaceFn(a, vo, so)

	fmt.Print("Enter time (t): ")
	_, err := fmt.Scan(&t)
	if err != nil {
		fmt.Println("Error reading time:", err)
		return
	}

	displacement := fn(t)
	fmt.Printf("Displacement after %.2f seconds: %.2f\n", t, displacement)
}
