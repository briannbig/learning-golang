package main

import "fmt"

func main() {
	var ke country
	ke.name = "Republic of Kenya"
	ke.capital = "Nairobi"
	ke.sports = append(ke.sports, "Athletes, football")

	usa := country{"United states of america", "Washington DC", sports{"American football, Basket Ball"}}

	nigeria := country{name: "Nigeria", capital: "lagos"}
	nigeria.sports = append(nigeria.sports, "Football")
	president := struct {
		name string; ctry country
	}{"Joe Biden", usa}

	printCountry(ke)
	printCountry(usa)
	printCountry(nigeria)

	fmt.Printf("president: %s ------> country: %s\n", president.name, president.ctry.name)
}

type sports []string
type country struct {
	name string
	capital string
	sports
}

func printCountry(myCountry country) {
	fmt.Printf("\nThe country is %s, and its capital is %s\n", myCountry.name, myCountry.capital)
	fmt.Println(myCountry.sports)
}