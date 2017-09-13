package main

import "fmt"

const (
	A = iota
	B
	C = 100
	D
	E = iota
)

func main() {
	fmt.Printf("A = %d\n", A)
	fmt.Printf("B = %d\n", B)
	fmt.Printf("C = %d\n", C)
	fmt.Printf("D = %d\n", D)
	fmt.Printf("E = %d\n", E)
}
