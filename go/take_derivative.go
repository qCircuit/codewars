package main

import (
	"fmt"
)

/*
func Derive(coefficient, exponent int) string {

	return strconv.Itoa(coefficient*exponent) + "x^" + strconv.Itoa(exponent-1)
}
*/

func Derive(coefficient, exponent int) string {
	return fmt.Sprintf("%dx^%d", coefficient*exponent, exponent-1)
}

func main() {
	fmt.Println(Derive(7, 8))
	fmt.Println(Derive(5, 9))
}
