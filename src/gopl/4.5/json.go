package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "星际穿越", Year: 2017, Color: true,
			Actors: []string{"a", "b"}},
		{Title: "详细信息", Year: 1995, Color: false,
			Actors: []string{"aa", "bb"}}}
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatal("JSON Marshal error:%s", err)
	}
	fmt.Printf("%s\n", data)
	data, err = json.MarshalIndent(movies, "", "   ")
	if err != nil {
		log.Fatal("JSON MarshalIndent error:%s", err)
	}
	fmt.Printf("%s\n", data)

	var titles []struct{ Title string }

	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatal("JSON Unmarshal error:%s", err)
	}
	fmt.Println(titles)
}
