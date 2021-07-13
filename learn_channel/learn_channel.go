package main

import (
	"fmt"
	"time"
)

func test1() {

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

func testclose() {
	var c chan string
	c = make(chan string, 3)
	c <- "hello"

	go func() {
		for {
			select {
			case a, ok := <-c:
				{
					if !ok {
						break
					}
					fmt.Println(a)
				}
			default:
			}
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("%d", i)
	}
	close(c)
	c <- "the world"

	time.Sleep(8 * time.Second)
}

func main() {
	testclose()
}
