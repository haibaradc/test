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
	var s2 = ""
	var s3 = ""
	index := strings.Index(s, ".")
	if index != -1 {
		s2 = s[index:]
		s3 = s[:index]
	} else {
		s3 = s
	}
	if s3[0] == '-' || s3[0] == '+' {
		s1 += s3[0:1]
		s3 = s3[1:]
	}
	l := len(s3)
	if l <= 3 {
		s1 += s3
		return s1
	}
	l2 := l % 3
	if l2 == 0 {
		l2 = 3
	}
	s1 += s3[0:l2]
	for i := l2; i < l; i += 3 {
		s1 += ","
		s1 += s3[i : i+3]
	}
	s1 += s2
	return s1
}

func main() {
	fmt.Println(basename("/a/b/c.d.go"))
	fmt.Println(basename2("/a/b/c.d.go"))
	fmt.Println(comma("+123464542332532432312312542543523.666666"))
}
