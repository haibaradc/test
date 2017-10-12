package main

import (
	"fmt"
	"time"
)

//30天之前
const thirtyDays = 30 * 24 * 60 * 60
const oneYears = 365 * 24 * 60 * 60

func main() {
	now := time.Now()
	td := time.Unix(now.Unix()-thirtyDays, 0)
	oy := time.Unix(now.Unix()-oneYears, 0)
	fmt.Println(td)
	fmt.Println(oy)
}
