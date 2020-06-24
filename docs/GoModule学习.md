在Goland中设置使用Go Modules

环境变量代理设置：
GOPROXY=https://goproxy.cn
export GOPROXY=https://goproxy.io


go help
go mod help
go mod：查看所有指令




go mod init 
go mod tidy



添加新依赖包
方法一：
直接修改 go.mod 文件，然后执行 go mod download
方法二：
使用 go get packagename@v1.2.3，会自动更新 go.mod 文件的
方法三：
go run、go build 也会自动下载依赖




参考
https://blog.csdn.net/e421083458/article/details/89762113
https://www.jianshu.com/p/dca7c631587f
https://zhuanlan.zhihu.com/p/116410261




