# Go Mac开发环境常见问题汇总

## zsh 转为 bash

`chsh -s /bin/bash`

设置完之后，要关闭terminal重新打开才能生效



## bash转为zsh

`chsh -s /bin/zsh`

设置完之后，要关闭terminal重新打开才能生效

`cat /etc/shells`可以看到Mac支持的所有shell



## .bash_profile设置

对于Go开发，经常需要使用`go install`下载到`$GOPATH/bin`下的命令，如果不把`$GOPATH/bin`放在环境变量里，在terminal里执行命令就会提示`command not found`。

可以把`GOROOT/bin`和`GOPATH/bin`都加入到环境变量设置里，参考如下：

```bash
export GOROOT=/usr/local/opt/go/libexec
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```



## 把文件内容拷贝到剪切板

* `cat file.txt | pbcopy`，在terminal里执行该命令

​	执行后，就可以使用CMD+V把剪切板里的内容拷贝到其它地方

* `pbpaste`：在terminal里执行该命令

​	执行后，可以把内容拷贝到终端上



## Mac安装homebrew

官网：[https://brew.sh/](https://brew.sh/)，在terminal里执行如下命令：

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```



## Mac安装MySQL

1. 使用homebrew安装

```bash
brew install mysql
```

2. 默认安装后root用户登录是不需要密码的，可以如下命令给root用户设置密码：

```bash
mysql_secure_installation
```

3. 升级后重启mysql

```bash
brew services restart mysql
```



## Mac安装Redis

1. 使用homebrew安装

   ```bash
   brew install redis
   ```

2. 升级后重启redis

   ```bash
   brew services restart redis
   ```

3. 关闭redis

   ```bash
   brew services stop redis
   ```

4. 启动redis

   ```bash
   brew services start redis
   ```


使用homebrew安装Redis后，相关文件的路径如下：

* `redis.conf`路径：/usr/local/etc/redis.conf
* `redis-server`路径：/usr/local/opt/redis/bin/redis-server
