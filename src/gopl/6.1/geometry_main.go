package main

import (
	"fmt"
	"gopl/6.1/geometry"
)

func main() {
	p1 := geometry.Point{1, 2}
	p2 := geometry.Point{4, 6}
	var p3 = geometry.Point{1, 2}
	var p4 geometry.Point
	p4.X = 4
	p4.Y = 6

	fmt.Println(geometry.Distance(p1, p2))
	fmt.Println(p1.Distance(p2))
	fmt.Println(p3.Distance(p4))

	path1 := geometry.Path{
		geometry.Point{1, 2},
		geometry.Point{4, 6},
		geometry.Point{8, 9},
	}
	fmt.Println(path1.Distance())

	p3.ScaleBy(2)
	fmt.Println(p3)
}
