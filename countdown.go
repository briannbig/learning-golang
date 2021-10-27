package main

import (
	"fmt"
	"time"
)

func main() {
	count := 10
	for ; count > 0; count-- {
		fmt.Println(count)
		time.Sleep(time.Second)
	}
	fmt.Println("Lights off")
}