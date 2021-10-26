package main

import "fmt"

var sentence string
var emptyString string = ""
var no, yes, maybe = "no", "yes", "maybe"

func main() {
	output()
}
func output() {
	fmt.Println(sentence)
	fmt.Println(emptyString)
	fmt.Println(no, yes, maybe)

}
