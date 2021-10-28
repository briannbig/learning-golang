package main

import (
	"fmt"
	"time"
)

func main() {
	for i:= 1; i <=5; i++{
		go wakeUp(3, i)
	}
	time.Sleep(15 * time.Second)
}
func wakeUp(counter int, id int) {
	fmt.Printf("%d ...waking.up...\n", id)
	for counter >= 0 {
		time.Sleep(time.Second)
		fmt.Printf("%d waking up in %d seconds...\n",id, counter)
		counter --
	}
	fmt.Printf("%d Finally up!!\n", id)
}
