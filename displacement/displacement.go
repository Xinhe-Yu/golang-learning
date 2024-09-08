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

	fmt.Print("Enter acceleration (a): ")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Error reading acceleration:", err)
		return
	}

	fmt.Print("Enter initial velocity (vo): ")
	_, err = fmt.Scan(&vo)
	if err != nil {
		fmt.Println("Error reading initial velocity:", err)
		return
	}

	fmt.Print("Enter initial displacement (so): ")
	_, err = fmt.Scan(&so)
	if err != nil {
		fmt.Println("Error reading initial displacement:", err)
		return
	}

	fn := GenDisplaceFn(a, vo, so)

	fmt.Print("Enter time (t): ")
	_, err = fmt.Scan(&t)
	if err != nil {
		fmt.Println("Error reading time:", err)
		return
	}

	displacement := fn(t)
	fmt.Printf("Displacement after %.2f seconds: %.2f\n", t, displacement)
}
