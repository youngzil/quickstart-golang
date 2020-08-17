package main

import (
	"context"
	"errors"
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
func GetRedisAddr(env string) (string, error) {
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
func CreateRedisClient(addr string) (*redis.ClusterClient, error) {
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
func GetApi(client *redis.ClusterClient, apiCode string) string {
	cmd := client.Get(ctx, "ApiOpen:"+apiCode)
	result, _ := cmd.Result()
	return result
}

// 获取App
func GetApp(client *redis.ClusterClient, apiCode string) string {
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
func GetVersions(client *redis.ClusterClient, apiCode string) []string {
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
func GetGroups(client *redis.ClusterClient, apiCode string) []string {
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
func GetParams(client *redis.ClusterClient, apiCode string) []string {
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
func GetNode(client *redis.ClusterClient, nodeCode string) string {
	cmd := client.Get(ctx, "ServiceNode:"+nodeCode)
	result, _ := cmd.Result()
	return result
}

// 美化
func Pretty(data string) string {
	//if data == "" {
	//	return ""
	//}
	//var str bytes.Buffer
	//_ = json.Indent(&str, []byte(data), "", "    ")
	//return str.String()
	return data
}
