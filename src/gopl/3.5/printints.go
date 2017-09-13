package main

import (
	"bytes"
	"fmt"
)

func intstostring(arr []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range arr {
		if i > 0 {
			buf.WriteByte(',')
		}
		fmt.Println(v)
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intstostring([]int{1, 2, 3}))
}
