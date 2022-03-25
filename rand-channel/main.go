package main

import (
	"fmt"
	"math/rand"
	"time"
)

func counter(c chan int) {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100)
	fmt.Printf("generated number: %d\n", num)
	c <- num
}
func main() {
	fmt.Println("Programm start")

	chan1 := make(chan int)
	chan2 := make(chan int)
	go counter(chan1)
	go counter(chan2)
	fmt.Println(<-chan1)
	fmt.Println(<-chan2)
}
