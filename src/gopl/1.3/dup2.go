package main

import (
	//"bufio"
	"bufio"
	"fmt"
	"os"
)

type item struct {
	count int
	is    map[string]bool
}

func countLines(f *os.File, counts map[string]item) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		tmp := counts[input.Text()] //map作为参数传递无法直接访问map中结构体的内部
		if tmp.count == 0 {
			tmp.is = make(map[string]bool)
		}
		tmp.is[f.Name()] = true
		tmp.count++
		counts[input.Text()] = tmp
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("命令行参数错误")
		return
	}
	counts := make(map[string]item)
	files := os.Args[1:]
	for _, file := range files {

		f, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		countLines(f, counts)
		f.Close()
	}

	for line, n := range counts {
		fmt.Printf("[%d]:%s  ", n.count, line)
		for isl, _ := range n.is {
			fmt.Printf("%s ", isl)
		}
		fmt.Printf("\n")
	}
}
