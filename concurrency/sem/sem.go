package main

import (
	"fmt"
	"sync"
	"time"
)

// c := [][]
// c := [goRoutine][goRoutine2]
// c := [goRoutine3][goRoutine2]

func main() {
	c := make(chan int, 6)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomething(i, &wg, c)
	}
	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("Id %d started\n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("Id %d Finished\n", i)
	<-c
}
