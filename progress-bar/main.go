package main

import (
	"sync"
	"time"

	pb "github.com/schollz/progressbar/v3"
)

func progressbar1(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	bar1 := pb.Default(100)
	for i := 0; i < 100; i++ {
		bar1.Add(1)
		time.Sleep(time.Duration(num) * time.Millisecond)
	}
}
func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go progressbar1(&wg, 100)
	wg.Wait()
}
