package main

import "fmt"

func clean(strings []string) []string {
	//for i, s := range strings {
	for i := 0; i < len(strings)-1; i++ {
		if strings[i+1] == strings[i] {
			copy(strings[i+1:], strings[i+2:])
			strings = strings[:len(strings)-1]
			i-- //下标不前进
		}
	}
	return strings
}

func clean2(strings []string) []string {
	//var nString []string
	var nString = []string{}
	fmt.Println(len(nString), cap(nString))
	for _, s := range strings {
		if len(nString) == 0 || s != nString[len(nString)-1] {
			nString = append(nString, s)
		}
	}
	return nString
}

func cleanSpace(strings []string) []string {
	//for i, s := range strings {
	for i := 0; i < len(strings)-1; i++ {
		if strings[i+1] == strings[i] && strings[i] == " " {
			copy(strings[i+1:], strings[i+2:])
			strings = strings[:len(strings)-1]
			i-- //下标不前进
		}
	}
	return strings
}

func main() {
	var strings = []string{"aa", "bb", "bb", "bb", "cc", "cc"}

	//strings = clean(strings)
	strings = clean2(strings)
	//strings = cleanSpace(strings)

	fmt.Println(strings)
}
