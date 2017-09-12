package main

import (
	"fmt"
	"gopl/2.5/ftoc_package"
	"gopl/2.5/mytools"
	"os"
	"strconv"
)

func main() {
	const freezingF, boilingF = 32.0, 212.0

	if len(os.Args) <= 1 {
		fmt.Println("参数错误")
		os.Exit(1)
	}
	fmt.Printf("%g°F = %g°C\n", freezingF, ftoc_package.FtoC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, ftoc_package.FtoC(boilingF))

	r, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("参数错误2")
		os.Exit(2)
	}
	fmt.Printf("半径：%g\n", r)
	fmt.Printf("直径：%g\n", mytools.GetL(r))
	fmt.Printf("周长：%g\n", mytools.GetC(r))
	fmt.Printf("面积：%g\n", mytools.GetS(r))
}
