package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("args error")
		return
	}
	counts := make(map[string]int)
	files, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	f_str := string(files)
	fmt.Println(f_str)
	lines := strings.Split(f_str, "\r\n")
	fmt.Println(lines)
	for _, line := range lines {
		counts[line]++
	}
	for line, n := range counts {
		fmt.Printf("[%d]:%s\n", n, line)
	}
}
