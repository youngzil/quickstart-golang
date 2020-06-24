package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strings"
)

var ctx = context.Background()

var xm = "10.1.243.23:7000"

//var xm = "127.0.0.1:6379"
//var xm = "20.26.37.179:28001"
var xyw = "20.26.103.195:6601"
var zfb = "20.26.103.210:6601"

// 获取redis地址
func getRedisAddr(env string) (string, error) {
	env = strings.TrimSpace(env)
	if env == "1" {
		return xm, nil
	} else if env == "2" {
		return xyw, nil
	} else if env == "3" {
		return zfb, nil
	} else {
		return "", errors.New("序号错误")
	}
}

// 创建Redis客户端
func createRedisClient(addr string) (*redis.ClusterClient, error) {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{addr},
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	} else {
		return client, nil
	}
}

// 获取Api
func getApi(client *redis.ClusterClient, apiCode string) string {
	cmd := client.Get(ctx, "ApiOpen:"+apiCode)
	result, _ := cmd.Result()
	return result
}

// 获取App
func getApp(client *redis.ClusterClient, apiCode string) string {
	key := ""
	_ = client.ForEachMaster(ctx, func(ctx context.Context, cli *redis.Client) error {
		cmd := cli.Keys(ctx, "ApiAppRelation*"+apiCode)
		result, _ := cmd.Result()
		if result != nil && len(result) > 0 {
			key = result[0]
		}
		return nil
	})
	if key == "" {
		return ""
	} else {
		cmd := client.Get(ctx, key)
		result, _ := cmd.Result()
		return result
	}
}

// 获取版本
func getVersions(client *redis.ClusterClient, apiCode string) []string {
	versions := make([]string, 0)
	_ = client.ForEachMaster(ctx, func(ctx context.Context, cli *redis.Client) error {
		cmd := cli.Keys(ctx, "ApiVersion:"+apiCode+":*")
		result, _ := cmd.Result()
		if result != nil && len(result) > 0 {
			versions = append(versions, result...)
		}
		return nil
	})
	list := make([]string, 0)
	if len(versions) == 0 {
		return list
	} else {
		for _, key := range versions {
			cmd := client.Get(ctx, key)
			result, _ := cmd.Result()
			list = append(list, result)
		}
	}
	return list
}

// 获取路由
func getGroups(client *redis.ClusterClient, apiCode string) []string {
	versions := make([]string, 0)
	_ = client.ForEachMaster(ctx, func(ctx context.Context, cli *redis.Client) error {
		cmd := cli.Keys(ctx, "ServiceGroup:"+apiCode+":*")
		result, _ := cmd.Result()
		if result != nil && len(result) > 0 {
			versions = append(versions, result...)
		}
		return nil
	})
	list := make([]string, 0)
	if len(versions) == 0 {
		return list
	} else {
		for _, key := range versions {
			cmd := client.Get(ctx, key)
			result, _ := cmd.Result()
			list = append(list, result)
		}
	}
	return list
}

// 获取版本
func getParams(client *redis.ClusterClient, apiCode string) []string {
	versions := make([]string, 0)
	_ = client.ForEachMaster(ctx, func(ctx context.Context, cli *redis.Client) error {
		cmd := cli.Keys(ctx, "ApiParam*"+apiCode+"*")
		result, _ := cmd.Result()
		if result != nil && len(result) > 0 {
			versions = append(versions, result...)
		}
		return nil
	})
	list := make([]string, 0)
	if len(versions) == 0 {
		return list
	} else {
		for _, key := range versions {
			cmd := client.Get(ctx, key)
			result, _ := cmd.Result()
			list = append(list, result)
		}
	}
	return list
}

// 获取节点
func getNode(client *redis.ClusterClient, nodeCode string) string {
	cmd := client.Get(ctx, "ServiceNode:"+nodeCode)
	result, _ := cmd.Result()
	return result
}

// 美化
func pretty(data string) string {
	//if data == "" {
	//	return ""
	//}
	//var str bytes.Buffer
	//_ = json.Indent(&str, []byte(data), "", "    ")
	//return str.String()
	return data
}

func main() {

	for true {

		fmt.Println("服务网关 Redis 配置检查工具（测试环境）\n")

		env := ""
		fmt.Print("环境列表:\n1.项目\n2.新业务\n3.准发布\n请输入环境序号: ")
		fmt.Scanln(&env)

		addr, err := getRedisAddr(env)
		if err != nil {
			fmt.Println("err: 序号错误")
			return
		}

		client, err := createRedisClient(addr)
		if err != nil {
			fmt.Println("err: " + err.Error())
			return
		}

		apiCode := ""
		fmt.Print("请输入 ApiCode: ")
		fmt.Scanln(&apiCode)
		apiCode = strings.TrimSpace(apiCode)

		api := getApi(client, apiCode)
		if api == "" {
			fmt.Println("Api 不存在")
			continue
		}

		fmt.Println("\n== 应用配置 ==")
		fmt.Println(pretty(getApp(client, apiCode)))

		fmt.Println("\n== 接口配置 ==")
		fmt.Println(pretty(api))

		fmt.Println("\n== 接口参数 ==")
		params := getParams(client, apiCode)
		for _, param := range params {
			fmt.Println(pretty(param))
		}

		fmt.Println("\n== 接口版本 ==")
		versions := getVersions(client, apiCode)
		for _, version := range versions {
			fmt.Println(pretty(version))
		}

		fmt.Println("\n== 路由分组 ==")
		groups := getGroups(client, apiCode)
		nodes := make([]string, 0)
		for _, group := range groups {
			fmt.Println(pretty(group))
			groupCode := group[strings.Index(group, "groupCode\":")+12 : strings.Index(group, "\",\"id\"")]
			nodes = append(nodes, getNode(client, groupCode))
		}

		fmt.Println("\n== 路由节点 ==")
		for _, node := range nodes {
			fmt.Println(pretty(node))
		}

		fmt.Println()
	}
}
