package main

import (
	"fmt"
	"math"
)

const pi int = math.Pi
func main() {
	rad := 7
	area := calculateArea(rad)
	fmt.Println(area)

}
func calculateArea(radius int) int {
	return pi * radius * radius
}
