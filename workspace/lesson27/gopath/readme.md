# Demo使用手册

1. 将package1这个文件夹拷贝到$GOPATH/src目录下

2. 在package1/main下，打开terminal

3. 在terminal里先设置关闭GO111MODULE

   ```sh
   export GO111MODULE=off
   ```

4. 在terminal里执行如下编译和运行命令

   ```go
   go build -o main main.go util.go // 会生成一个名为main的可执行文件
   ./main // 运行可执行文件
   ```

   

