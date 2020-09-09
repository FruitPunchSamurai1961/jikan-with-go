package jikan

import "fmt"

func main() {
	data, err := getPics(1)
	fmt.Println(data.Pictures[0].Small)
	fmt.Println(err)
}
