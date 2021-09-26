package main

import (
	"fmt"
	"time"

	"github.com/swordhell/etcdutil"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	cfg := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2381"},
		DialTimeout: 5 * time.Second,
	}
	if util, err := etcdutil.NewETCDUtil(cfg, 5, nil); err == nil {

	} else {
		fmt.Println(err.Error())
	}
}
