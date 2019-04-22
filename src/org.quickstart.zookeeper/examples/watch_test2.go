package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

//部分代码如下：
func main() {
	option := zk.WithEventCallback(callback)

	conn, _, err := zk.Connect(hosts, time.Second*5, option)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	_, _, ech, err := conn.ExistsW(path1)
	if err != nil {
		fmt.Println(err)
		return
	}

	create(conn, path1, data1)

	go watchCreataNode(ech)
}

func watchCreataNode(ech <-chan zk.Event) {
	event := <-ech
	fmt.Println("*******************")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
	fmt.Println("-------------------")
}

//输出如下：
/*******************
path: /whatyy
type: EventNodeCreated
state: Unknown
-------------------
*/
