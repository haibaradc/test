package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, v := range strings {
		if v != "" {
			strings[i] = v
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	var new []string
	for _, v := range strings {
		if v != "" {
			new = append(new, v)
		}
	}
	return new
}

func main() {
	//strings := []string{"a", "", "c"}
	//var strings = []string{"a", "", "c"}
	//var strings []string = []string{"a", "", "c"}
	var strings = make([]string, 3)
	//copy(strings, []string{"a", "", "c"})
	strings = append(strings, "a")
	strings = append(strings, "")
	strings = append(strings, "c")
	strings = nonempty2(strings)
	fmt.Println(strings)
}
