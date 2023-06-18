package main

import "fmt"

func Hero(bullets, dragons int) bool {
	return bullets >= dragons*2
}

func main() {
	fmt.Println(Hero(5, 15))
}
