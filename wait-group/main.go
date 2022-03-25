package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello from goroutine")
	time.Sleep(time.Second * 1)
}

func hello(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Hello from %d \n", id)

}

func main() {
	fmt.Println("Promgramm starts")

	// trace.Start(os.Stderr)
	// defer trace.Stop()

	var wg1 sync.WaitGroup

	wg1.Add(1)
	go sayHello(&wg1)
	wg1.Wait()
	fmt.Println("Hello from main")

	fmt.Println("New greeting programm starts")
	const numGreaters = 5
	var wg2 sync.WaitGroup
	wg2.Add(numGreaters)
	for i := 0; i < numGreaters; i++ {
		go hello(&wg2, i)
	}
	wg2.Wait()

}
