package main

import (
	"fmt"
	"time"
)

func main() {
	var stringStream chan string
	stringStream = make(chan string, 1)
	var recvStream <-chan string
	var sendStream chan<- string

	recvStream = stringStream
	sendStream = stringStream

	go func() {
		time.Sleep(2 * time.Second)
		for {
			if val, ok := <-recvStream; ok {
				fmt.Println(val)
			} else {
				fmt.Println("no data")
				break
			}
		}
	}()
	sendStream <- "hh"
	sendStream <- "hh"
	close(stringStream)
	time.Sleep(8 * time.Second)
}
