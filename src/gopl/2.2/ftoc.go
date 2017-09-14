package main

import "fmt"

func ftoc(f float64) float64 {
	return (f - 32) * 5 / 9
}

func f() *int {
	v := 1
	return &v
}

func main() {
	var freezingF, boilingF = 32.0, 212.0

	fmt.Printf("%g°F = %g°C\n", freezingF, ftoc(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, ftoc(boilingF))

	fmt.Println(f())
	fmt.Println(f())
}
