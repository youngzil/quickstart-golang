go build hello.go
./hello
或者
go run hello.go

Golang依赖管理工具dep
https://github.com/golang/dep


文档
https://golang.github.io/dep/


安装
Linux curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
MacOS
$ brew install dep
$ brew upgrade dep

安装后
dep version


$GOPATH/src路径下的项目的根目录下执行 dep init就会下载依赖到$GOPATH/pkg/dep/sources下面





1、第一行代码 package main 定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
2、下一行 import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。
3、下一行 func main() 是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。
4、当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。
5、以 // 开头的单行注释。多行注释也叫块注释，均已以 /* 开头，并以 */ 结尾，且不可以嵌套使用，多行注释一般用于包的文档描述或注释成块的代码片段。
6、下一行 fmt.Println(...) 可以将字符串输出到控制台，并在最后自动增加换行字符 \n。 使用 fmt.Print("hello, world\n") 可以得到相同的结果。 
7、需要注意的是 { 不能单独放在一行
8、在 Go 程序中，一行代表一个语句结束。将多个语句写在同一行，它们则必须使用 ; 人为区分，但在实际开发中我们并不鼓励这种做法。
9、Go 语言中变量的声明必须使用空格隔开
10、定义变量的4种方式
11、
12、
13、
14、
15、
16、




-------------------------------------------------------------------------------------------------
1、golang的命名需要使用驼峰命名法，且不能出现下划线

2、golang中根据首字母的大小写来确定可以访问的权限。无论是方法名、常量、变量名还是结构体的名称，如果首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用
  可以简单的理解成，首字母大写是公有的，首字母小写是私有的
  
3、结构体中属性名的大写
如果属性名小写则在数据解析（如json解析,或将结构体作为请求或访问参数）时无法解析
 --------------------------------------------------------------------------------------------------------

Golang文件名命名规则
文件名_平台.go
文件名_test.go或者 文件名_平台_test.go
文件名_版本号.go
文件名_(平台:可选)_CPU类型.go


注意点一、
go build 的时候会选择性地编译以系统名结尾的文件(linux、darwin、windows、freebsd)。例如Linux(Unix)系统下编译只会选择array_linux.go文件，其它系统命名后缀文件全部忽略。

 注意点二、
在xxx.go文件的文件头上添加 // + build !windows (tags)，可以选择在windows系统下面不编译 
// +build !windows
package main

总结：golang跨平台没有java好用，但是跟c语言差不多，都得针对不同平台不同特性迭轮子





文件命名规范
https://golang.google.cn/doc/effective_go.html#mixed-caps
https://golang.org/doc/effective_go.html#package-names
https://www.cnblogs.com/hetonghai/p/9049536.html




https://golang.org
https://github.com/golang/go


中文社区
https://studygolang.com/
https://gocn.vip/
https://github.com/shen100/golang123


学习示例
https://github.com/SimonWaldherr/golang-examples




学习教程
http://www.runoob.com/go/go-tutorial.html
https://blog.csdn.net/han0373/article/category/7676220
https://github.com/Unknwon/the-way-to-go_ZH_CN
https://github.com/dariubs/GoBooks



