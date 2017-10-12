package main

import "fmt"

func main() {
	var runes []rune

	for index, item := range "Hello 世界!" {
		fmt.Println(index)
		runes = append(runes, item)
	}
	fmt.Printf("%q\n", runes)
}
