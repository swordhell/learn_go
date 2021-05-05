package main

import (
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	for {
		time.Sleep(0)
	}
	wg.Wait()
}
