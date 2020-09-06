package main

import "fmt"

func main() {
	data, err := GetAnimeById(10)
	fmt.Println(data)
	fmt.Println(err)
}
