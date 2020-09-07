package jikan

import "fmt"

func main() {
	data, err := getNews(1)
	fmt.Println(data)
	fmt.Println(err)
}
