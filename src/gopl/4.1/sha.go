package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var flags = flag.Int("n", 256, "256/384/512")

func main() {
	flag.Parse()
	fmt.Printf("%d\n", *flags)
	if *flags == 256 {
		s1 := sha256.Sum256([]byte("x"))
		fmt.Printf("SHA256:%x\n", s1)
	} else if *flags == 384 {
		s1 := sha512.Sum384([]byte("x"))
		fmt.Printf("SHA384:%x\n", s1)
	} else if *flags == 512 {
		s1 := sha512.Sum512([]byte("x"))
		fmt.Printf("SHA512:%x\n", s1)
	} else {
		fmt.Println("参数错误")
	}
}
