package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep, item string
	for i := 0; i < len(os.Args); i++ {
		item = os.Args[i]
		s += sep + os.Args[i]
		sep = " "
		fmt.Println(item)
	}
	fmt.Println(s)
}
