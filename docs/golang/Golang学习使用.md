基本语法






参考
http://c.biancheng.net/view/91.html


---------------------------------------------------------------------------------------------------------------------  
基本语法

访问一个包中的函数的语法是 <package>.Function()，变量 <package>.Var。
在 Go 中，当变量或函数的首字母大写的时候，函数会被从包中导出（在包外部可见，或者说公有的）

注1：包和package有关系：main包和main函数比较特殊,是程序的入口，其他的包就要和文件夹名字一样
注2：go 里面一个目录为一个package, 一个package级别的func, type, 变量, 常量, 这个package下的所有文件里的代码都可以随意访问, 也不需要首字母大写.

概括来说：
公有函数/变量的名字以大写字母开头；
私有函数/变量的名字以小写字母开头。






