package main

import "fmt"

func ftoc(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, ftoc(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, ftoc(boilingF))
}
