package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if line == "exit" {
			break
		}
		counts[line]++
	}
	for line, n := range counts {
		fmt.Printf("[%d]:%s\n", n, line)
	}
	a, b := counts["abc"]
	fmt.Println(a, b)
}
