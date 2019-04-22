//如下：
package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

var hosts = []string{"localhost:2181"}

var path1 = "/whatzk"

var flags int32 = zk.FlagEphemeral
var data1 = []byte("hello,this is a zk go test demo!!!")
var acls = zk.WorldACL(zk.PermAll)

func main() {
	option := zk.WithEventCallback(callback)

	conn, _, err := zk.Connect(hosts, time.Second*5, option)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	_, _, _, err = conn.ExistsW(path1)
	if err != nil {
		fmt.Println(err)
		return
	}

	create(conn, path1, data1)

	time.Sleep(time.Second * 2)

	_, _, _, err = conn.ExistsW(path1)
	if err != nil {
		fmt.Println(err)
		return
	}
	delete(conn, path1)

}

func callback(event zk.Event) {
	fmt.Println("*******************")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
	fmt.Println("-------------------")
}

func create(conn *zk.Conn, path string, data []byte) {
	_, err_create := conn.Create(path, data, flags, acls)
	if err_create != nil {
		fmt.Println(err_create)
		return
	}

}

func delete(conn *zk.Conn, path string) {
	err_delete := conn.Delete(path, -1)
	if err_delete != nil && err_delete != zk.ErrNoNode {
		fmt.Println("Delete returned error: %+v", err_delete)
		return
	}
	fmt.Println("delete success:")
}

//输出:
/*******************
path:
type: EventSession
state: StateConnecting
-------------------
*******************
path:
type: EventSession
state: StateConnected
-------------------
*******************
path:
type: EventSession
state: StateHasSession
-------------------
*******************
path: /whatzk
type: EventNodeCreated
state: Unknown
-------------------
*******************
path: /whatzk
type: EventNodeDeleted
state: Unknown
-------------------
*/
