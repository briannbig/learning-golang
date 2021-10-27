package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	degrees := 0
	for {
		fmt.Println(degrees)
		time.Sleep(time.Millisecond*100)
		degrees++
		if degrees >= 360 {
			degrees = 0
			if rand.Intn(2) == 0 {
				break
			}
		}
	}

}
