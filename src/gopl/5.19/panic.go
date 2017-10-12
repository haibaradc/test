package main

import "fmt"

func panic_test(i int) {
	defer func() {
		r := recover()
		if r == nil {
			fmt.Println("no error")
		} else {
			fmt.Println(r)
		}
	}()
	panic(i)
}

func main() {
	panic_test(1)

}
