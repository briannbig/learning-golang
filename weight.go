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
	var grm grams = 10000
	fmt.Println(grm.gramsToKilograms()) // calling gramsToKilogramsMethod
	fmt.Println(grm.multiplyGrams(3))
}
//function with custom types
func gainWeight(initial kilograms, gained kilograms) kilograms {
	return initial + gained
}
func weightToGrams(weight kilograms) grams{
	return grams(weight * 1000)
}

//method belonging to grams, converts grams to kilograms, takes no parameter
func (g grams) gramsToKilograms() kilograms {
	return kilograms(g / 1000)
}
//method belonging to grams, converts grams to kilograms, takes one int parameter
func (g grams) multiplyGrams(number int) grams {
	return grams(number) * g
}

