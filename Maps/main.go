package main

import (
	"fmt"
)

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Hex color of:", color, "is", hex)

	}
}

func main() {
	//colors := map[string]string{
	//	"red":  "#ff80000",
	//	"blue": "#ff90000",
	//}
	//fmt.Printf("%+v\n", colors)
	//fmt.Println(colors)

	//colors := make(map[string]string)
	//colors["name"] = "Thomas"
	//fmt.Println(colors)
	//delete(colors, "name")
	//fmt.Println(colors)

	//colors := map[string]string{
	//	"red":  "#ff80000",
	//	"blue": "#ff90000",
	//}
	//printMap(colors)

	

}
