package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	guess := 1
	var random = rand.Intn(time.Now().Second())
	ans := guess == random
	fmt.Println(random)
	fmt.Println(ans)
}
