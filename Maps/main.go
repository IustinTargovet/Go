package main

import (
	"fmt"
)

func printMap(c map[int]string) {
	for i, color := range c {
		if i == 1 {
			fmt.Println("first color is", color)
		}
		if i == 2 {
			fmt.Println("second color is", color)
		}
		if i == 3 {
			fmt.Println("third color is", color)
		}
	}
}

func main() {
	colors := map[int]string{
		1: "red",
		2: "green",
		3: "blue",
	}

	printMap(colors)
}
