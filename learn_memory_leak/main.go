package main

import (
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func fun1(wg *sync.WaitGroup, memories [][]byte) {
	defer wg.Done()
	for {
		time.Sleep(time.Second * 1)
		tmp := make([]byte, 6*1024*1024)
		memories = append(memories, tmp)
	}
}

func main() {
	wg := &sync.WaitGroup{}

	var memories [][]byte
	go func() {
		http.ListenAndServe(":10003", nil)
	}()
	wg.Add(1)

	go func() {
		fun1(wg, memories)
	}()

	wg.Wait()
}
