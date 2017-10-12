package main

import "fmt"

func square(a int) int {
	return a * a
}

func negative(a int) int {
	return -a
}

func negative2(a int, f func(a int) int) int {
	return f(a)
}

func product(a int, b int) int {
	return a * b
}

func product2(a, b, c int) int {
	return a * b * c
}

func main() {
	f := square

	fmt.Println(f(5))

	f = negative

	fmt.Println(f(5))

	f1 := negative2

	fmt.Println(f1(5, negative))

	f2 := product

	fmt.Println(f2(2, 3))

	f3 := product2

	fmt.Println(f3(2, 3, 4))
}
