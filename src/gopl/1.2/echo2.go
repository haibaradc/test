package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	for _, item := range os.Args[:] {
		s += sep + item
		sep = " "
		fmt.Println(item)
	}
	fmt.Println(s)
	fmt.Println(strings.Join(os.Args[:], " "))
	fmt.Println(os.Args[:])
}
