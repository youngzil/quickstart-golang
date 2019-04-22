package main

import (
	//etcd_client "go.etcd.io/etcd/clientv3"
	"context"
	//"encoding/json"
	"fmt"
	"time"
	//"github.com/coreos/etcd/clientv3"
	"go.etcd.io/etcd/clientv3"
)

var (
	endpoints      = []string{"localhost:2379", "localhost:22379", "localhost:32379"}
	dialTimeout    = 5 * time.Second
	requestTimeout = 10 * time.Second
)

func main() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	fmt.Println("connect succ")
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err = cli.Put(ctx, "/logagent/conf/", "sample_value")
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), requestTimeout)
	resp, err := cli.Get(ctx, "/logagent/conf/")
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}

}
