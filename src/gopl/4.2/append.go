package main

import "fmt"

func myAppend(list []int, item int) []int {
	var newList []int
	c := cap(list)
	l := len(list)
	if l >= c {
		newList = make([]int, l+1, c*2)
		copy(newList, list)
	} else {
		newList = list[:]
	}
	newList[l] = item
	return newList
}

func main() {
	list := []int{1, 2, 3}
	fmt.Println(len(list))
	fmt.Println(cap(list))
	list2 := myAppend(list, 5)
	fmt.Println(list2)
	fmt.Println(len(list2))
	fmt.Println(cap(list2))
	list2 = append(list2, 9)
	fmt.Println(list2)

	var list3 []int
	//list3 = make([]int, 1)
	list3 = append(list3, 5)
	fmt.Println(list3)
}
