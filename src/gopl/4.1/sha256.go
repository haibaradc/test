package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		var count byte = 0
		for j := 0; j < 8; j++ {
			if (i>>uint(j))&1 == 1 {
				count++
			}
		}
		pc[i] = count
	}

}

func popCount(value [32]byte) int {
	var count = 0
	for _, v := range value {
		count += int(pc[int(v)])
	}
	return count
}

func main() {
	s1 := sha256.Sum256([]byte("x"))
	s2 := sha256.Sum256([]byte("X"))
	var s3 [32]byte
	fmt.Printf("s1 = %x\n", s1)
	fmt.Printf("s2 = %x\n", s2)
	/*
		for i, v := range []byte("12345") {
			fmt.Printf("%d:%c\n", i, v)
		}
	*/
	fmt.Println(popCount(s1))
	fmt.Println(popCount(s2))
	fmt.Println(popCount(s3))
}
