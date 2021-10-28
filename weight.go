package main

import "fmt"
// custom type kilograms
type kilograms int
type grams int
func main() {
	var weight kilograms = 50
	var gainedWeight = 10
	//weight += gainedWeight // Invalid operation: mismatched types kilograms and int
	weight += kilograms(gainedWeight) //conversion to kilograms type needed first
	fmt.Println(weight)
	fmt.Println(gainWeight(50, 20))
	fmt.Println(weightToGrams(kilograms(80)))// casting 80 to kilograms
	fmt.Println(weightToGrams(70))// compiler will automatically detect the int as kilograms type
}
//function with custom types
func gainWeight(initial kilograms, gained kilograms) kilograms {
	return initial + gained
}
func weightToGrams(weight kilograms) grams{
	return grams(weight * 1000)
}

