package main

/*
import (
	"fmt"
	"strconv"
	"strings"
)

func Points(games []string) int {
	var totalPoints int

	for _, value := range games {
		game := strings.Split(value, ":")
		a, _ := strconv.Atoi(game[0])
		b, _ := strconv.Atoi(game[1])

		if a > b {
			totalPoints += 3
		} else if a == b {
			totalPoints += 1
		}
	}

	return totalPoints
}
*/

import (
	"fmt"
	"strings"
)

func Points(games []string) int {
	totalPoints := 0
	for _, value := range games {
		game := strings.Split(value, ":")
		x, y := game[0], game[1]
		switch {
		case x > y:
			totalPoints += 3
		case x == y:
			totalPoints += 1
		}
	}
	return totalPoints

}

func main() {
	games := []string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"}
	fmt.Println(Points(games))
}
