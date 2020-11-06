package main

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/sirupsen/logrus"
)

func trans1(waitSec int) {
	c1, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer c1.Close()
	time.Sleep(time.Duration(1) * time.Second)
	c1.Do("WATCH lock lock_times")
	defer c1.Do("UNWATCH")
	c1.Do("MULTI")
	c1.Do(`SET lock "joe"`)
	c1.Do(`INCR lock_times`)
	time.Sleep(time.Duration(waitSec) * time.Second)
	if reply, err := c1.Do(`EXEC`); err != nil {
		logrus.Warn(waitSec, " ", err.Error())
	} else {
		logrus.Info(waitSec, " ", reply)
	}

}

func main() {
	go trans1(3)
	go trans1(5)

	time.Sleep(time.Duration(10) * time.Second)
	logrus.Debug("main connect success.")
}
