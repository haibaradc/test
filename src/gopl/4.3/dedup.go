package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]uint)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		seen[line]++
		fmt.Println(line, seen[line])
	}
	if err := input.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
