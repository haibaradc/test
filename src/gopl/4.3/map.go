package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := make(map[string]int)
	ages["tom"] = 25
	ages["bob"] = 27
	ages["kit"] = 30

	ages["bob"]++

	delete(ages, "kit")
	//ages2 := map[string]int{"bob2": 20, "tom2": 25}

	for k, v := range ages {
		fmt.Println(k, v)
	}
	fmt.Println()

	var names []string
	for k := range ages {
		names = append(names, k)
	}
	sort.Strings(names)
	for _, v := range names {
		fmt.Println(v, ages[v])
	}
}
