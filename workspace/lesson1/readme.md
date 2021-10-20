# Go的程序结构
* 包声明
* 引入包
* 函数
* 变量
* 语句和表达式
* 注释

# 注意事项
* func main()是程序开始执行的函数(但是如果有func init()函数，则会先执行init函数，再执行main函数)

# 编译和运行
Go是编译型语言
* 编译+运行分步执行 
    * go build hello.go
    * ./hello
* 编译+运行一步到位
    * go run hello.go 