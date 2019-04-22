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

func ExampleWatcher_watchWithPrefix() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})

	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	defer cli.Close()

	rch := cli.Watch(context.Background(), "foo", clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}

	wc := cli.Watch(context.Background(), "/job/", clientv3.WithPrefix(), clientv3.WithPrevKV())
	for v := range wc {
		if v.Err() != nil {
			panic(err)
		}
		for _, e := range v.Events {
			fmt.Printf("type:%v\n kv:%v  prevKey:%v  ", e.Type, e.Kv, e.PrevKv)
		}
	}

	// PUT "foo1" : "bar"
}
