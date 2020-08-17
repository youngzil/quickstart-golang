package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

func main() {

	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}

	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超时的连接就close
	DB.SetMaxOpenConns(100)                  //设置最大连接数
	CreateTable(DB)
	InsertData(DB)
	QueryOne(DB)
	QueryMulti(DB)
	UpdateData(DB)
	DeleteData(DB)

	for true {

		fmt.Println("服务网关 Redis 配置检查工具（测试环境）\n")

		env := ""
		fmt.Print("环境列表:\n1.项目\n2.新业务\n3.准发布\n请输入环境序号: ")
		fmt.Scanln(&env)

		addr, err := GetRedisAddr(env)
		if err != nil {
			fmt.Println("err: 序号错误")
			return
		}

		client, err := CreateRedisClient(addr)
		if err != nil {
			fmt.Println("err: " + err.Error())
			return
		}

		apiCode := ""
		fmt.Print("请输入 ApiCode: ")
		fmt.Scanln(&apiCode)
		apiCode = strings.TrimSpace(apiCode)

		api := GetApi(client, apiCode)
		if api == "" {
			fmt.Println("Api 不存在")
			continue
		}

		fmt.Println("\n== 应用配置 ==")
		fmt.Println(Pretty(GetApp(client, apiCode)))

		fmt.Println("\n== 接口配置 ==")
		fmt.Println(Pretty(api))

		fmt.Println("\n== 接口参数 ==")
		params := GetParams(client, apiCode)
		for _, param := range params {
			fmt.Println(Pretty(param))
		}

		fmt.Println("\n== 接口版本 ==")
		versions := GetVersions(client, apiCode)
		for _, version := range versions {
			fmt.Println(Pretty(version))
		}

		fmt.Println("\n== 路由分组 ==")
		groups := GetGroups(client, apiCode)
		nodes := make([]string, 0)
		for _, group := range groups {
			fmt.Println(Pretty(group))
			groupCode := group[strings.Index(group, "groupCode\":")+12 : strings.Index(group, "\",\"id\"")]
			nodes = append(nodes, GetNode(client, groupCode))
		}

		fmt.Println("\n== 路由节点 ==")
		for _, node := range nodes {
			fmt.Println(Pretty(node))
		}

		fmt.Println()
	}
}
