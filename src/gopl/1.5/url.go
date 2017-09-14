package main

import (
	"fmt"
	"io"
	//"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "参数错误\n")
		os.Exit(-1)
	}
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error:%v\n", err)
			os.Exit(-1)
		}
		fmt.Println(resp.Status)
		fmt.Println(resp.StatusCode)
		//b, err := ioutil.ReadAll(resp.Body)
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error:%v\n", err)
			os.Exit(-1)
		}
		//fmt.Println(string(b))
	}
}
