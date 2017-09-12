package main

import "fmt"

const boilingF = 3.14

func main() {
	var f = boilingF
	var c = f*35 - 16
	fmt.Printf("f = %f,c = %f\n", f, c)
}
