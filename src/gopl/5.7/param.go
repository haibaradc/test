package main

import (
	"fmt"
	"strings"
)

func max(n ...int) (int, bool) {
	var maxI = 0
	if len(n) == 0 {
		return maxI, false
	} else {
		maxI = n[0]
	}
	for _, v := range n {
		if v > maxI {
			maxI = v
		}
	}
	return maxI, true
}

func min(n ...int) (int, bool) {
	var minI = 0
	if len(n) == 0 {
		return minI, false
	} else {
		minI = n[0]
	}
	for _, v := range n {
		if v < minI {
			minI = v
		}
	}
	return minI, true
}

func myjoin(s string, i ...string) string {
	var items string
	for _, item := range i {
		items += item
	}
	return strings.Replace(s, " ", items, -1)
}

func main() {
	fmt.Println(max())
	fmt.Println(min(-1, -55))
	fmt.Println(myjoin("123 456 789", ",", "?"))
}
