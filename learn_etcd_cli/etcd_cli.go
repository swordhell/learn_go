package main

import (
	"fmt"
	"os"
	"time"

	"github.com/swordhell/etcdutil"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	cfg := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2381"},
		DialTimeout: 5 * time.Second,
	}
	c := make(chan os.Signal, 1)
	if util, err := etcdutil.NewETCDUtil(cfg, 5, nil); err == nil {

		if handle, err := util.WatchKey("serverlist", true, func(eventType etcdutil.EventType, key string, value string) {
			fmt.Println("eventType ", eventType, " key ", key, " value ", value)
		}); err == nil {
			fmt.Println("WatchKey handle ", handle)
		} else {
			fmt.Println("WatchKey fail ", err.Error())
		}
		<-c
	} else {
		fmt.Println(err.Error())
	}
}
