package main

import "fmt"

var pc [256]byte
var pc2 [256]byte

func init() {
	for i := range pc {
		var count byte = 0
		for j := 0; j < 8; j++ {
			if (i>>uint(j))&1 == 1 {
				count++
			}
		}
		pc2[i] = count
		pc[i] = pc[i/2] + byte(i&1)
	}

}

// 1 2 4 8
func PopCount(i uint64) int {
	return int(pc[byte(i>>(0*8))] +
		pc[byte(i>>(1*8))] +
		pc[byte(i>>(2*8))] +
		pc[byte(i>>(3*8))] +
		pc[byte(i>>(4*8))] +
		pc[byte(i>>(5*8))] +
		pc[byte(i>>(6*8))] +
		pc[byte(i>>(7*8))])
}

func main() {
	a := PopCount(15)
	fmt.Println(pc)
	fmt.Println(pc2)
	fmt.Println(a)
}
