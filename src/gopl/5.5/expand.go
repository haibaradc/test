package main

import "fmt"

func foo(s string) string {
	if s == "foo" {
		return "FOO"
	} else {
		return s
	}
}

func expand(s string, f func(s string) string) string {
	if s == "foo" {
		return f(s)
	} else {
		return s
	}
}

func main() {
	fmt.Println(expand("abc", foo))
	fmt.Println(expand("foo", foo))
}
