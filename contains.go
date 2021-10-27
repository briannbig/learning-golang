package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Find yourself a dimly lit cavern.")
	var command = "Walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Println("you leave the cave", exit)
}
