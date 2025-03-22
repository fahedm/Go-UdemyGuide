package main

import "fmt"

func main() {
	colors := make(map[string]string)
	// colors := map[string]string{
	// 	"red":"#ff0000",
	// 	"green":"#000000",
	// }

	colors["white"] = "#ffffff"
	fmt.Println(colors)
}