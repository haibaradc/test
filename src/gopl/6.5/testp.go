package main

import "fmt"

type TestP struct {
	words []uint64
}

func main() {
	var x *TestP = new(TestP)

	x.words = make([]uint64, 1)
	x.words[0] = 1

	var y TestP = *x //为什么改了y的值，x的值也变了呢，是因为words是动态分配的，多个结构体变量对应同一个words，就
	//类似于c的malloc，多个变量对应的都是一块内存

	fmt.Println("before	x:", x.words)
	fmt.Println("before	y:", y.words)
	y.words[0] = 0
	fmt.Println("after	x:", x.words)
	fmt.Println("after	y:", y.words)

}
