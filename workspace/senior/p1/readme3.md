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
PATH="/Library/Frameworks/Python.framework/Versions/3.9/bin:${PATH}"
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

