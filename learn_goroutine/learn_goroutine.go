package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("hello")

	done := make(chan int)
	linkData := make(chan string, 3)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {

			select {
			case <-done:
				fmt.Println("done")
				return
			case node, ok := <-linkData:
				if !ok {
					fmt.Println("exit by !ok")
					return
				}
				fmt.Println("node ", node)
				time.Sleep(5 * time.Second)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	linkData <- "hello"
	time.Sleep(2 * time.Second)
	linkData <- "abel"
	time.Sleep(3 * time.Second)
	close(linkData)

	wg.Wait()
}
