package main

import (
	"fmt"
	"sync"
)

type A struct {
	Age int32
}

func main() {
	c := make(chan *A)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(c chan *A) {
		defer wg.Done()
		for {
			select {
			case p, ok := <-c:
				if ok {
					fmt.Println(p)
				} else {
					fmt.Println("read chan fail")
				}
			}
		}
	}(c)

	var a *A

	c <- a

	wg.Wait()
}
