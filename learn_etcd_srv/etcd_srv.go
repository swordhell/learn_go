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

		if err := util.PutWithLease("serverlist/", "{state:0}"); err != nil {
			fmt.Println(err.Error())
		}
		<-c
	} else {
		fmt.Println(err.Error())
	}
}
