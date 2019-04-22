package main

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {

	multiConn()

	//连接到ZK server端
	var hosts = []string{"localhost:2181"} //server端host
	zkClient, _, err_connect := zk.Connect(hosts, time.Second*5)
	if err_connect != nil {
		fmt.Println(err_connect)
		return
	}
	defer zkClient.Close()

	//节点增删改查
	var path = "/test"
	var data = []byte("hello zk")
	var flags int32 = 0
	//flags有4种取值：
	//0:永久，除非手动删除
	//zk.FlagEphemeral = 1:短暂，session断开则改节点也被删除
	//zk.FlagSequence  = 2:会自动在节点后面添加序号
	//3:Ephemeral和Sequence，即，短暂且自动添加序号
	var acls = zk.WorldACL(zk.PermAll) //控制访问权限模式

	p, err_create := zkClient.Create(path, data, flags, acls)
	if err_create != nil {
		fmt.Println(err_create)
		return
	}
	fmt.Println("create:", p) //输出 create：/test

	//get获取数据
	var getdata, _, err_get = zkClient.Get(path)
	if err_get != nil {
		fmt.Println("Get failed on node 2: %+v", err_get)
		return
	}
	fmt.Println("get:", string(getdata))

	//设置数据
	statData, err_set := zkClient.Set(path, []byte("hello zk2"), -1)
	if err_set != nil {
		fmt.Println("set failed on node 2: %+v", err_set)
		return
	}
	fmt.Println("set Stat data:", statData)

	//get获取数据
	getdata, _, err_get = zkClient.Get(path)
	if err_get != nil {
		fmt.Println("after set ，Get failed on node 2: %+v", err_get)
		return
	}
	fmt.Println("get:", string(getdata))

	//ls查看子节点
	childPath, statData2, error_child := zkClient.Children("/")
	if error_child != nil {
		fmt.Println("children returned error: %+v", error_child)
		return
	}
	fmt.Println("children:", childPath, statData2)

	//删除数据
	err_delete := zkClient.Delete(path, -1)
	if err_delete != nil && err_delete != zk.ErrNoNode {
		fmt.Println("Delete returned error: %+v", err_delete)
		return
	}
	fmt.Println("delete success:")

}

func multiConn() {

	/*在主机的/etc/hosts中配置host
	127.0.0.1 host1
	127.0.0.1 host2
	127.0.0.1 host3*/

	//使用步骤如下：（相关代码位于dnshostprovider.go中）

	var hosts = []string{"host1:2181", "host2:2181", "host3:2181"}
	hostPro := new(zk.DNSHostProvider)
	err := hostPro.Init(hosts) //先初始化
	if err != nil {
		fmt.Println(err)
		return
	}
	server, retryStart := hostPro.Next() //获得host

	fmt.Println("conn success:", server, retryStart)

	hostPro.Connected() //连接成功后会调用
	//上面的一系列步骤都集成在func Connect(servers []string, sessionTimeout time.Duration, options ...connOption) (*Conn, <-chan Event, error)中
}
