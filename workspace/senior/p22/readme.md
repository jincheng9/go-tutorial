# 官方教程：Go fuzzing入门

## 前言

Go 1.18在go工具链里引入了fuzzing模糊测试，可以帮助我们发现Go代码里的漏洞或者可能导致程序崩溃的输入。Go官方团队也在官网发布了fuzzing入门教程，帮助大家快速上手。

![](../../img/fuzzing.png)

本人对Go官方教程在翻译的基础上做了一些表述上的优化，以飨读者。

**注意**：fuzzing模糊测试和Go已有的单元测试以及性能测试框架是互为补充的，并不是替代关系。



## 教程内容

这篇教程会介绍Go fuzzing的入门基础知识。fuzzing可以构造随机数据来找出代码里的漏洞或者可能导致程序崩溃的输入。通过fuzzing可以找出的漏洞包括SQL注入、缓存溢出、拒绝服务(Denial of Service)攻击和XSS(cross-site scripting)攻击等。

在这个教程里，你会给一个函数写一段fuzz test(模糊测试)程序，然后运行go命令来发现代码里的问题，最后通过调试来修复问题。

本文里涉及的专业术语，可以参考 [Go Fuzzing glossary](https://go.dev/doc/fuzz/#glossary)。

接下来会按照如下章节介绍：

1. [为你的代码创建一个目录](#为你的代码创建一个目录)
2. [实现一个函数](#实现一个函数)
3. [增加单元测试](#增加单元测试)
4. [增加模糊测试](#增加模糊测试)
5. [修复2个bug](#修复2个bug)
6. [总结](#总结)

## 准备工作

* 安装Go 1.18 Beta 1或者更新的版本。安装指引可以参考[下面的介绍](#安装和使用beta版本)。
* 有一个代码编辑工具。任何文本编辑器都可以。
* 有一个命令行终端。Go可以运行在Linux，Mac上的任何命令行终端，也可以运行在Windows的PowerShell或者cmd之上。
* 有一个支持fuzzing的环境。目前Go fuzzing只支持AMD64和ARM64架构。

### 安装和使用beta版本

这个教程需要使用Go 1.18 Beta 1版本了的泛型功能。使用如下步骤，安装beta版本 

1. 使用下面的命令安装beta版本

   ```sh
   $ go install golang.org/dl/go1.18beta1@latest
   ```

2. 运行如下命令来下载更新

   ```sh
   $ go1.18beta1 download
   ```

   **注意**：如果在MAC或者Linux上执行`go1.18beta1`提示`command not found`，需要设置`bash`或者`zsh`对应的profile环境变量文件。`bash`设置在`~/.bash_profile`文件里，内容为：

   ```bash
   export GOROOT=/usr/local/opt/go/libexec
   export GOPATH=$HOME/go
   export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
   ```

   `GOROOT`和`GOPATH`的值可以通过`go env`命令查看，设置完后执行`source ~/.bash_profile`让设置生效，再执行`go1.18beta1`就不报错了。

3. 使用beta版本的go命令，不要去使用release版本的go命令

   你可以通过直接使用`go1.18beta1`命令或者给`go1.18beta1`起一个简单的别名

   * 直接使用`go1.18beta1`命令

     ```sh
     $ go1.18beta1 version
     ```

   * 给`go1.18beta1`命令起一个别名

     ```sh
     $ alias go=go1.18beta1
     $ go version
     ```

   下面的教程都假设你已经把`go1.18beta1`命令设置了别名`go`。



##   为你的代码创建一个目录

首先创建一个目录用于存放你写的代码。

1. 打开一个命令行终端，切换到你的`home`目录

   * 在Linux或者Mac上执行如下命令(Linux或者Mac上只需要执行`cd`就可以进入到`home`目录)

     ```bash
     cd
     ```

   * 在Windows上执行如下命令

     ```po
     C:\> cd %HOMEPATH%
     ```

2. 在命令行终端，创建一个名为`fuzz`的目录，并进入该目录

   ```bas
   $ mkdir fuzz
   $ cd fuzz
   ```

3. 创建一个go module

   运行`go mod init`命令，来给你的项目设置module路径

   ```bash
   $ go mod init example/fuzz
   ```

   **注意**：对于生产代码，你可以根据项目实际情况来指定module路径，如果想了解更多，可以参考[Go Module依赖管理](https://go.dev/doc/modules/managing-dependencies)。

接下来，我们来使用map写一些简单的代码来做字符串的反转，然后使用fuzzing来做模糊测试。

## 实现一个函数

在这个章节，你需要实现一个函数来对字符串做反转。

### 编写代码

1. 打开你的文本编辑器，在fuzz目录下创建一个`main.go`源文件。

2. 在`main.go`里编写如下代码：

   ```go
   // maing.go
   package main
   
   import "fmt"
   
   func Reverse(s string) string {
   	b := []byte(s)
   	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
   		b[i], b[j] = b[j], b[i]
   	}
   	return string(b)
   }
   
   func main() {
   	input := "The quick brown fox jumped over the lazy dog"
   	rev := Reverse(input)
   	doubleRev := Reverse(rev)
   	fmt.Printf("original: %q\n", input)
   	fmt.Printf("reversed: %q\n", rev)
   	fmt.Printf("reversed again: %q\n", doubleRev)
   }
   ```

### 运行代码

在`main.go`所在目录执行命令`go run .`来运行代码，结果如下：

```bash
$ go run .
original: "The quick brown fox jumped over the lazy dog"
reversed: "god yzal eht revo depmuj xof nworb kciuq ehT"
reversed again: "The quick brown fox jumped over the lazy dog"
```

## 增加单元测试

在这个章节，你会给`Reverse`函数编写单元测试代码。

### 编写单元测试

1. 在fuzz目录下创建文件`reverse_test.go`。

2. 在`reverse_test.go`里编写如下代码：

   ```go
   package main
   
   import (
       "testing"
   )
   
   func TestReverse(t *testing.T) {
       testcases := []struct {
           in, want string
       }{
           {"Hello, world", "dlrow ,olleH"},
           {" ", " "},
           {"!12345", "54321!"},
       }
       for _, tc := range testcases {
           rev := Reverse(tc.in)
           if rev != tc.want {
                   t.Errorf("Reverse: %q, want %q", rev, tc.want)
           }
       }
   }
   ```

### 运行单元测试

使用`go test`命令来运行单元测试

```
$ go test
PASS
ok      example/fuzz  0.013s
```

接下来，我们给`Reverse`函数增加模糊测试(fuzz test)代码。

## 增加模糊测试

单元测试有局限性，每个测试输入必须由开发者指定加到单元测试的测试用例里。

fuzzing的优点之一是可以基于开发者代码里指定的测试输入作为基础数据，进一步自动生成新的随机测试数据，用来发现指定测试输入没有覆盖到的边界情况。

在这个章节，我们会把单元测试转换成模糊测试，这样可以更轻松地生成更多的测试输入。

**注意：**你可以把单元测试、性能测试和模糊测试放在同一个`*_test.go`文件里。

### 编写模糊测试

在文本编辑器里把`reverse_test.go`里的单元测试代码`TestReverse`替换成如下的模糊测试代码`FuzzReverse`。

```go
func FuzzReverse(f *testing.F) {
    testcases := []string{"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc)  // Use f.Add to provide a seed corpus
    }
    f.Fuzz(func(t *testing.T, orig string) {
        rev := Reverse(orig)
        doubleRev := Reverse(rev)
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}
```

Fuzzing也有一定的局限性。

在单元测试里，因为测试输入是固定的，你可以知道调用`Reverse`函数后每个输入字符串得到的反转字符串应该是什么，然后在单元测试的代码里判断`Reverse`的执行结果是否和预期相符。例如，对于测试用例`Reverse("Hello, world")`，单元测试预期的结果是 `"dlrow ,olleH"`。

但是使用fuzzing时，我们没办法预期输出结果是什么，因为测试的输入除了我们代码里指定的用例之外，还有fuzzing随机生成的。对于随机生成的测试输入，我们当然没办法提前知道输出结果是什么。

虽然如此，本文里的`Reverse`函数有几个特性我们还是可以在模糊测试里做验证。

1. 对一个字符串做2次反转，得到的结果和源字符串相同
2. 反转后的字符串也仍然是一个有效的UTF-8编码的字符串

**注意**：fuzzing模糊测试和Go已有的单元测试以及性能测试框架是互为补充的，并不是替代关系。

比如我们实现的`Reverse`函数如果是一个错误的版本，直接return返回输入的字符串，是完全可以通过上面的模糊测试的，但是没法通过我们前面编写的单元测试。因此单元测试和模糊测试是互为补充的，不是替代关系。

Go模糊测试和单元测试在语法上有如下差异：

- Go模糊测试函数以FuzzXxx开头，单元测试函数以TestXxx开头

- Go模糊测试函数以 `*testing.F`作为入参，单元测试函数以`*testing.T`作为入参

- Go模糊测试会调用`f.Add`函数和`f.Fuzz`函数。

  - `f.Add`函数把指定输入作为模糊测试的种子语料库(seed corpus)，fuzzing基于种子语料库生成随机输入。
  - `f.Fuzz`函数接收一个fuzz target函数作为入参。fuzz target函数有多个参数，第一个参数是`*testing.T`，其它参数是被模糊的类型(**注意:**被模糊的类型目前只支持部分内置类型, 列在 [Go Fuzzing docs](https://go.dev/doc/fuzz/#requirements)，未来会支持更多的内置类型)。

  ![](../../img/fuzzing.png)

上面的`FuzzReverse`函数里用到了`utf8`这个package，因此要在`reverse_test.go`开头import这个package，参考如下代码：

```
package main

import (
    "testing"
    "unicode/utf8"
)
```

### 运行模糊测试

1. 执行如下命令来运行模糊测试。

   这个方式只会使用种子语料库，而不会生成随机测试数据。通过这种方式可以用来验证种子语料库的测试数据是否可以测试通过。

   ```bash
   $ go test
   PASS
   ok      example/fuzz  0.013s
   ```

   如果`reverse_test.go`文件里有其它单元测试函数或者模糊测试函数，但是只想运行`FuzzReverse`模糊测试函数，我们可以执行`go test -run=FuzzReverse`命令。

   **注意**：`go test`默认会执行所有以`TestXxx`开头的单元测试函数和以`FuzzXxx`开头的模糊测试函数，默认不运行以`BenchmarkXxx`开头的性能测试函数，如果我们想运行 benchmark用例，则需要加上 `-bench` 参数。

2. 如果要基于种子语料库生成随机测试数据用于测试，需要给`go test`命令增加` -fuzz`参数，使用方法如下：

   ```
   $ go test -fuzz=Fuzz
   fuzz: elapsed: 0s, gathering baseline coverage: 0/3 completed
   fuzz: elapsed: 0s, gathering baseline coverage: 3/3 completed, now fuzzing with 8 workers
   fuzz: minimizing 38-byte failing input file...
   --- FAIL: FuzzReverse (0.01s)
       --- FAIL: FuzzReverse (0.00s)
           reverse_test.go:20: Reverse produced invalid UTF-8 string "\x9c\xdd"
   
       Failing input written to testdata/fuzz/FuzzReverse/af69258a12129d6cbba438df5d5f25ba0ec050461c116f777e77ea7c9a0d217a
       To re-run:
       go test -run=FuzzReverse/af69258a12129d6cbba438df5d5f25ba0ec050461c116f777e77ea7c9a0d217a
   FAIL
   exit status 1
   FAIL    example/fuzz  0.030s
   ```

   上面的fuzzing测试结果是`FAIL`，引起`FAIL`的输入数据被写到了一个种子语料库文件里。下次运行`go test`命令的时候，即使没有`-fuzz`参数，这个种子语料库文件里的输入也会被用到。

   可以用文本编辑器打开`testdata/fuzz/FuzzReverse`目录下的文件，看看引起Fuzzing测试失败的测试数据长什么样。下面是一个示例文件，你那边运行后得到的测试数据可能和这个不一样，但文件里的内容格式会是一样的。

   ```
   go test fuzz v1
   string("泃")
   ```

   种子语料库文件里的第一行标识的是编码版本(说直白点，就是这个种子语料库文件里的内容格式)，虽然目前只有v1这1个版本，但是Fuzzing设计者考虑到未来可能引入新的编码版本，于是加了一个编码版本的概念。

   接下来的每一行是具体的测试数据。因为本文的

   The first line of the corpus file indicates the encoding version. Each following line represents the value of each type making up the corpus entry. Since the fuzz target only takes 1 input, there is only 1 value after the version.

3. Run `go test` again without the` -fuzz` flag; the new failing seed corpus entry will be used:

   ```
   $ go test
   --- FAIL: FuzzReverse (0.00s)
       --- FAIL: FuzzReverse/af69258a12129d6cbba438df5d5f25ba0ec050461c116f777e77ea7c9a0d217a (0.00s)
           reverse_test.go:20: Reverse produced invalid string
   FAIL
   exit status 1
   FAIL    example/fuzz  0.016s
   ```

   既然Go fuzzing测试没通过，那就需要我们调试代码来找出问题所在了。

## 修复2个bug

在这个章节，我们会调试程序，修复Go fuzzing测出来的bug。

你可以自己花一些时间思考下，先尝试自己解决问题。

### 定位问题

There are a few different ways you could debug this error. If you are using VS Code as your text editor, you can [set up your debugger](https://github.com/golang/vscode-go/blob/master/docs/debugging.md) to investigate.

In this tutorial, we will log useful debugging info to your terminal.

First, consider the docs for [`utf8.ValidString`](https://pkg.go.dev/unicode/utf8).

```
ValidString reports whether s consists entirely of valid UTF-8-encoded runes.
```

The current `Reverse` function reverses the string byte-by-byte, and therein lies our problem. In order to preserve the UTF-8-encoded runes of the original string, we must instead reverse the string rune-by-rune.

To examine why the input (in this case, the Chinese character `泃`) is causing `Reverse` to produce an invalid string when reversed, you can inspect the number of runes in the reversed string.

#### 编写代码

In your text editor, replace the fuzz target within `FuzzReverse` with the following.

```
f.Fuzz(func(t *testing.T, orig string) {
    rev := Reverse(orig)
    doubleRev := Reverse(rev)
    t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
    if orig != doubleRev {
        t.Errorf("Before: %q, after: %q", orig, doubleRev)
    }
    if utf8.ValidString(orig) && !utf8.ValidString(rev) {
        t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
    }
})
```

This `t.Logf` line will print to the command line if an error occurs, or if executing the test with `-v`, which can help you debug this particular issue.

#### 运行代码

Run the test using go test

```
$ go test
--- FAIL: FuzzReverse (0.00s)
    --- FAIL: FuzzReverse/28f36ef487f23e6c7a81ebdaa9feffe2f2b02b4cddaa6252e87f69863046a5e0 (0.00s)
        reverse_test.go:16: Number of runes: orig=1, rev=3, doubleRev=1
        reverse_test.go:21: Reverse produced invalid UTF-8 string "\x83\xb3\xe6"
FAIL
exit status 1
FAIL    example/fuzz    0.598s
```

The entire seed corpus used strings in which every character was a single byte. However, characters such as 泃 can require several bytes. Thus, reversing the string byte-by-byte will invalidate multi-byte characters.

**Note:** If you’re curious about how Go deals with strings, read the blog post [Strings, bytes, runes and characters in Go](https://go.dev/blog/strings) for a deeper understanding.

With a better understanding of the bug, correct the error in the `Reverse` function.

### 修复问题

To correct the `Reverse` function, let’s traverse the string by runes, instead of by bytes.

#### Write the code

In your text editor, replace the existing Reverse() function with the following.

```
func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
```

The key difference is that `Reverse` is now iterating over each `rune` in the string, rather than each `byte`.

#### Run the code

1. Run the test using `go test`

   ```
   $ go test
   PASS
   ok      example/fuzz  0.016s
   ```

   The test now passes!

2. Fuzz it again with `go test -fuzz`, to see if there are any new bugs.

   ```
   $ go test -fuzz=Fuzz
   fuzz: elapsed: 0s, gathering baseline coverage: 0/37 completed
   fuzz: minimizing 506-byte failing input file...
   fuzz: elapsed: 0s, gathering baseline coverage: 5/37 completed
   --- FAIL: FuzzReverse (0.02s)
       --- FAIL: FuzzReverse (0.00s)
           reverse_test.go:33: Before: "\x91", after: "�"
   
       Failing input written to testdata/fuzz/FuzzReverse/1ffc28f7538e29d79fce69fef20ce5ea72648529a9ca10bea392bcff28cd015c
       To re-run:
       go test -run=FuzzReverse/1ffc28f7538e29d79fce69fef20ce5ea72648529a9ca10bea392bcff28cd015c
   FAIL
   exit status 1
   FAIL    example/fuzz  0.032s
   ```

   We can see that the string is different from the original after being reversed twice. This time the input itself is invalid unicode. How is this possible if we’re fuzzing with strings?

   Let’s debug again.

## Fix the double reverse error

In this section, you will debug the double reverse failure and fix the bug.

Feel free to spend some time thinking about this and trying to fix the issue yourself before moving on.

### Diagnose the error

Like before, there are several ways you could debug this failure. In this case, using a [debugger](https://github.com/golang/vscode-go/blob/master/docs/debugging.md) would be a great approach.

In this tutorial, we will log useful debugging info in the `Reverse` function.

Look closely at the reversed string to spot the error. In Go, [a string is a read only slice of bytes](https://go.dev/blog/strings), and can contain bytes that aren’t valid UTF-8. The original string is a byte slice with one byte, `'\x91'`. When the input string is set to `[]rune`, Go encodes the byte slice to UTF-8, and replaces the byte with the UTF-8 character �. When we compare the replacement UTF-8 character to the input byte slice, they are clearly not equal.

#### Write the code

1. In your text editor, replace the `Reverse` function with the following.

   ```
   func Reverse(s string) string {
       fmt.Printf("input: %q\n", s)
       r := []rune(s)
       fmt.Printf("runes: %q\n", r)
       for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
           r[i], r[j] = r[j], r[i]
       }
       return string(r)
   }
   ```

   This will help us understand what is going wrong when converting the string to a slice of runes.

#### Run the code

This time, we only want to run the failing test in order to inspect the logs. To do this, we will use `go test -run`.

```
$ go test -run=FuzzReverse/28f36ef487f23e6c7a81ebdaa9feffe2f2b02b4cddaa6252e87f69863046a5e0
input: "\x91"
runes: ['�']
input: "�"
runes: ['�']
--- FAIL: FuzzReverse (0.00s)
    --- FAIL: FuzzReverse/28f36ef487f23e6c7a81ebdaa9feffe2f2b02b4cddaa6252e87f69863046a5e0 (0.00s)
        reverse_test.go:16: Number of runes: orig=1, rev=1, doubleRev=1
        reverse_test.go:18: Before: "\x91", after: "�"
FAIL
exit status 1
FAIL    example/fuzz    0.145s
```

To run a specific corpus entry within FuzzXxx/testdata, you can provide {FuzzTestName}/{filename} to `-run`. This can be helpful when debugging.

Knowing that the input is invalid unicode, let’s fix the error in our `Reverse` function.

### Fix the error

To fix this issue, let’s return an error if the input to `Reverse` isn’t valid UTF-8.

#### Write the code

1. In your text editor, replace the existing `Reverse` function with the following.

   ```
   func Reverse(s string) (string, error) {
       if !utf8.ValidString(s) {
           return s, errors.New("input is not valid UTF-8")
       }
       r := []rune(s)
       for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
           r[i], r[j] = r[j], r[i]
       }
       return string(r), nil
   }
   ```

   This change will return an error if the input string contains characters which are not valid UTF-8.

2. Since the Reverse function now returns an error, modify the `main` function to discard the extra error value. Replace the existing `main` function with the following.

   ```
   func main() {
       input := "The quick brown fox jumped over the lazy dog"
       rev, revErr := Reverse(input)
       doubleRev, doubleRevErr := Reverse(rev)
       fmt.Printf("original: %q\n", input)
       fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
       fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
   }
   ```

   These calls to `Reverse` should return a nil error, since the input string is valid UTF-8.

3. You will need to import the errors and the unicode/utf8 packages. The import statement in main.go should look like the following.

   ```
   import (
       "errors"
       "fmt"
       "unicode/utf8"
   )
   ```

4. Modify the reverse_test.go file to check for errors and skip the test if errors are generated by returning.

   ```
   func FuzzReverse(f *testing.F) {
       testcases := []string {"Hello, world", " ", "!12345"}
       for _, tc := range testcases {
           f.Add(tc)  // Use f.Add to provide a seed corpus
       }
       f.Fuzz(func(t *testing.T, orig string) {
           rev, err1 := Reverse(orig)
           if err1 != nil {
               return
           }
           doubleRev, err2 := Reverse(rev)
           if err2 != nil {
                return
           }
           if orig != doubleRev {
               t.Errorf("Before: %q, after: %q", orig, doubleRev)
           }
           if utf8.ValidString(orig) && !utf8.ValidString(rev) {
               t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
           }
       })
   }
   ```

   Rather than returning, you can also call `t.Skip()` to stop the execution of that fuzz input.

#### Run the code

1. Run the test using go test

   ```
   $ go test
   PASS
   ok      example/fuzz  0.019s
   ```

2. Fuzz it with `go test -fuzz=Fuzz`, then after a few seconds has passed, stop fuzzing with `ctrl-C`.

   ```
   $ go test -fuzz=Fuzz
   fuzz: elapsed: 0s, gathering baseline coverage: 0/38 completed
   fuzz: elapsed: 0s, gathering baseline coverage: 38/38 completed, now fuzzing with 4 workers
   fuzz: elapsed: 3s, execs: 86342 (28778/sec), new interesting: 2 (total: 35)
   fuzz: elapsed: 6s, execs: 193490 (35714/sec), new interesting: 4 (total: 37)
   fuzz: elapsed: 9s, execs: 304390 (36961/sec), new interesting: 4 (total: 37)
   ...
   fuzz: elapsed: 3m45s, execs: 7246222 (32357/sec), new interesting: 8 (total: 41)
   ^Cfuzz: elapsed: 3m48s, execs: 7335316 (31648/sec), new interesting: 8 (total: 41)
   PASS
   ok      example/fuzz  228.000s
   ```

   The fuzz test will run until it encounters a failing input unless you pass the `-fuzztime` flag. The default is to run forever if no failures occur, and the process can be interrupted with `ctrl-C`.

3. Fuzz it with `go test -fuzz=Fuzz -fuzztime 30s` which will fuzz for 30 seconds before exiting if no failure was found.

   ```
   $ go test -fuzz=Fuzz -fuzztime 30s
   fuzz: elapsed: 0s, gathering baseline coverage: 0/5 completed
   fuzz: elapsed: 0s, gathering baseline coverage: 5/5 completed, now fuzzing with 4 workers
   fuzz: elapsed: 3s, execs: 80290 (26763/sec), new interesting: 12 (total: 12)
   fuzz: elapsed: 6s, execs: 210803 (43501/sec), new interesting: 14 (total: 14)
   fuzz: elapsed: 9s, execs: 292882 (27360/sec), new interesting: 14 (total: 14)
   fuzz: elapsed: 12s, execs: 371872 (26329/sec), new interesting: 14 (total: 14)
   fuzz: elapsed: 15s, execs: 517169 (48433/sec), new interesting: 15 (total: 15)
   fuzz: elapsed: 18s, execs: 663276 (48699/sec), new interesting: 15 (total: 15)
   fuzz: elapsed: 21s, execs: 771698 (36143/sec), new interesting: 15 (total: 15)
   fuzz: elapsed: 24s, execs: 924768 (50990/sec), new interesting: 16 (total: 16)
   fuzz: elapsed: 27s, execs: 1082025 (52427/sec), new interesting: 17 (total: 17)
   fuzz: elapsed: 30s, execs: 1172817 (30281/sec), new interesting: 17 (total: 17)
   fuzz: elapsed: 31s, execs: 1172817 (0/sec), new interesting: 17 (total: 17)
   PASS
   ok      example/fuzz  31.025s
   ```

   Fuzzing passed!

   In addition to the `-fuzz` flag, several new flags have been added to `go test` and can be viewed in the [documentation](https://go.dev/doc/fuzz/#custom-settings).

## 总结

Nicely done! You’ve just introduced yourself to fuzzing in Go.

The next step is to choose a function in your code that you’d like to fuzz, and try it out! If fuzzing finds a bug in your code, consider adding it to the [trophy case](https://github.com/golang/go/wiki/Fuzzing-trophy-case).

If you experience any problems or have an idea for a feature, [file an issue](https://github.com/golang/go/issues/new/?&labels=fuzz).

For discussion and general feedback about the feature, you can also participate in the [#fuzzing channel](https://gophers.slack.com/archives/CH5KV1AKE) in Gophers Slack.

Check out the documentation at [go.dev/doc/fuzz](https://go.dev/doc/fuzz/#requirements) for further reading.

本文的完整代码参考：



## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## References

* https://go.dev/doc/tutorial/fuzz
* https://github.com/golang/go/issues/44551
* https://go.dev/doc/fuzz/