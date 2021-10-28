package main

import (
	"fmt"
	"time"
)

func main() {
	go wakeUp(10)
	time.Sleep(15 * time.Second)
}
func wakeUp(counter int) {
	fmt.Println("...waking.up...")
	for counter >= 0 {
		time.Sleep(time.Second)
		fmt.Printf("waking up in %d seconds...\n", counter)
		counter --
	}
	fmt.Println("Finally up!!")
}
