package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	cHTTP string = "http://"
)

func get_url(myurl string, ch chan string) {
	start := time.Now()
	if !strings.HasPrefix(myurl, cHTTP) {
		myurl = cHTTP + myurl
	}
	fmt.Println(myurl)
	resp, err := http.Get(myurl)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	num, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	seconds := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs,%d,%s", seconds, num, myurl)
}

func main() {
	ch := make(chan string)
	start := time.Now()

	for _, url := range os.Args[1:] {
		go get_url(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed!\n", time.Since(start).Seconds())
}
