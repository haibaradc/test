package main

import (
	"fmt"
	"os"
)

func main() {
	for i, item := range os.Args[:] {
		fmt.Println("[", i, "]", item)
	}
}
