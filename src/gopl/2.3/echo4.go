package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "输出后不换行")
var s = flag.String("s", "", "分隔符")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *s))
	if *n != true {
		fmt.Println()
	}
}
