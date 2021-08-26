Golang安装

安装
brew install go

brew info go
go version
go env

编辑.bash_profile文件，添加path

#golang安装
export GOROOT="/usr/local/Cellar/go/1.14.4/libexec"
export GOROOT=$HOME/software/go
export GOPATH=$HOME/gopath:$HOME/git/quickstart-golang
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
#(国外服务器请忽略此行输入)
export GOPROXY=https://goproxy.cn,direct




Linux安装：
1、解压tar
2、配置环境变量
export GOROOT=/home/aibuild45/yang/go
export GOPATH=/home/aibuild45/yang/go
export PATH=$PATH:$GOROOT/bin
3、go env

source .bash_profile
source ~/.zshrc

[Linux下安装配置Go语言环境](https://9yu.cc/index.php/archives/12/)



安装完成后
go version
go env  #查看安装信息



go build hello.go
./hello
或者
go run hello.go


