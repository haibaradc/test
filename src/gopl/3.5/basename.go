package main

import (
	"fmt"
	"strings"
)

// /a/b/c.d.go - d

func basename(s string) string {
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	f1 := strings.LastIndex(s, "/")
	s = s[f1+1:]
	f2 := strings.LastIndex(s, ".")
	if f2 != -1 {
		s = s[:f2]
	} else {

	}
	return s
}

// 12,345
func comma(s string) string {
	var s1 = ""
	l := len(s)
	if l <= 3 {
		s1 = s
		return s1
	}
	l2 := l % 3
	if l2 == 0 {
		l2 = 3
	}
	s1 += s[0:l2]
	for i := l2; i < l; i += 3 {
		s1 += ","
		s1 += s[i : i+3]
	}
	return s1
}

func main() {
	fmt.Println(basename("/a/b/c.d.go"))
	fmt.Println(basename2("/a/b/c.d.go"))
	fmt.Println(comma("123464542332532432312312542543523666666"))
}
