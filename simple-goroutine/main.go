package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main function")

	go countNumbers(20)

	fmt.Println("End main function")
	time.Sleep(2 * time.Second)
}

func countNumbers(limit int) {
	num := 0
	for i := 1; i < limit; i++ {
		num += i
	}
	fmt.Println("Num: ", num)
}
