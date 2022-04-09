# 官方博文：Go工作区模式最佳实践

## 前言

Go 1.18除了引入泛型(generics)、模糊测试(Fuzzing)之外，另外一个重大功能是引入了工作区模式(workspace mode)。

Go官方团队的*Beth Brown*于2022.04.05在官网上专门写了一篇博文，详细介绍了workspace模式的使用场景和最佳实践。

本人针对官方原文做了一个翻译，以飨读者。**同时在本文最后，附上了对workspace模式的入门介绍。**



## 原文翻译

Go官方团队*Beth Brown*

2022.04.05

Go 1.18新增了工作区模式(workspace mode)，让你可以同时跨多个Go Module进行开发。

你可以从[download](https://go.dev/dl/)地址下载Go 1.18，[release notes](https://go.dev/doc/go1.18)有更多关于Go 1.18的变化细节。



## 工作区(workspaces)

Go 1.18引入的[工作区](https://go.dev/ref/mod#workspaces)模式，可以让你不用修改每个Go Module的`go.mod`，就能同时跨多个Go Module进行开发。工作区里的每个Go Module在解析依赖的时候都被当做根Module。

在Go 1.18以前，如果遇到以下场景：`Module A新增了一个feature，Module B需要使用Module A的这个新feature`，你有2种方案：

* 发布Module A的修改到代码仓库，Module B更新依赖的Module A的版本即可
* 修改Module B的`go.mod`，使用`replace`指令把对Module A的依赖指向你本地未发布的Module A所在目录。等Module A发布后，在发布Module B的时候，再删除Module B的`go.mod`文件里的`replace`指令。

有了Go工作区模式之后，针对上述场景，我们有了更为简单的方案：你可以在工作区目录维护一个`go.work`文件来管理你的所有依赖。`go.work`里的`use`和`replace`指令会覆盖工作区目录下的每个Go Module的`go.mod`文件，因此没有必要去修改Go Module的`go.mod`文件了。



#### go work init

你可以使用`go work init`来创建一个workspace，`go work init` 的语法如下所示：

```bash
go work init [moddirs]
```

`moddirs`是Go Module所在的本地目录。如果有多个Go Module，就用空格分开。如果`go work init`后面没有参数，会创建一个空的workspace。

执行`go work init`后会生成一个`go.work`文件，`go.work`里列出了该workspace需要用到的Go Module，workspace目录不需要包含你当前正在开发的Go Module。



#### go work use

如果要给workspace新增Go Module，可以使用如下命令：

```bash
go work use [-r] moddir
```

或者手动编辑`go work`文件。

如果带有`-r`参数，会递归查找`-r`后面的路径参数下的所有子目录，把所有包含`go.mod`文件的子目录都添加到`go work`文件中。

如果某个目录已经被加到`go.work`里了，后面该目录没有`go.mod`文件了或者该目录被删除了，那对该目录再次执行`go work use`命令，该目录的`use`指令会从`go.work`文件里自动移除。(**注意**：自动移除要从Go 1.18正式版本才会生效，Go 1.18beta1版本不会生效)



#### go.work

`go.work`的语法和`go.mod`类似，包含如下3个指令：

- `go`: go的版本，例如 `go 1.18`
- `use`: adds a module on disk to the set of main modules in a workspace. Its argument is a relative path to the directory containing the module’s `go.mod` file. A `use` directive doesn’t add modules in subdirectories of the specified directory.
- `replace`: Similar to a `replace` directive in a `go.mod` file, a `replace` directive in a `go.work` file replaces the contents of a *specific version* of a module, or *all versions* of a module, with contents found elsewhere.



## 最佳实践

Workspace使用起来很灵活，接下来我们会介绍最常见的几种最佳实践。

### Add a feature to an upstream module and use it in your own module

1. Create a directory for your workspace.

2. Clone the upstream module you want to edit. If you haven’t contributed to Go before, read the [contribution guide](https://go.dev/doc/contribute).

3. Add your feature to the local version of the upstream module.

4. Run `go work init [path-to-upstream-mod-dir]` in the workspace folder.

5. Make changes to your own module in order to implement the feature added to the upstream module.

6. Run `go work use [path-to-your-module]` in the workspace folder.

   The `go work use` command adds the path to your module to your `go.work` file:

   ```
   go 1.18
   
   use (
          ./path-to-upstream-mod-dir
          ./path-to-your-module
   )
   ```

7. Run and test your module using the new feature added to the upstream module.

8. Publish the upstream module with the new feature.

9. Publish your module using the new feature.

### Work with multiple interdependent modules in the same repository

While working on multiple modules in the same repository, the `go.work` file defines the workspace instead of using `replace` directives in each module’s `go.mod` file.

1. Create a directory for your workspace.

2. Clone the repository with the modules you want to edit. The modules don’t have to be in your workspace folder as you specify the relative path to each with the `use` directive.

3. Run `go work init [path-to-module-one] [path-to-module-two]` in your workspace directory.

   Example: You are working on `example.com/x/tools/groundhog` which depends on other packages in the `example.com/x/tools` module.

   You clone the repository and then run `go work init tools tools/groundhog` in your workspace folder.

   The contents of your `go.work` file resemble the following:

   ```
   go 1.18
   
   use (
           ./tools
           ./tools/groundhog
   )
   ```

   Any local changes made in the `tools` module will be used by `tools/groundhog` in your workspace.

### Switching between dependency configurations

To test your modules with different dependency configurations you can either create multiple workspaces with separate `go.work` files, or keep one workspace and comment out the `use` directives you don’t want in a single `go.work` file.

To create multiple workspaces:

1. Create separate directories for different dependency needs.
2. Run `go work init` in each of your workspace directories.
3. Add the dependencies you want within each directory via `go work use [path-to-dependency]`.
4. Run `go run [path-to-your-module]` in each workspace directory to use the dependencies specified by its `go.work` file.

To test out different dependencies within the same workspace, open the `go.work` file and add or comment out the desired dependencies.

### Still using GOPATH?

Maybe using workspaces will change your mind. `GOPATH` users can resolve their dependencies using a `go.work` file located at the base of their `GOPATH` directory. Workspaces don’t aim to completely recreate all `GOPATH` workflows, but they can create a setup that shares some of the convenience of `GOPATH` while still providing the benefits of modules.

To create a workspace for GOPATH:

1. Run `go work init` in the root of your `GOPATH` directory.
2. To use a local module or specific version as a dependency in your workspace, run `go work use [path-to-module]`.
3. To replace existing dependencies in your modules' `go.mod` files use `go work replace [path-to-module]`.
4. To add all the modules in your GOPATH or any directory, run `go work use -r` to recursively add directories with a `go.mod` file to your workspace. If a directory doesn’t have a `go.mod` file, or no longer exists, the `use` directive for that directory is removed from your `go.work` file.

> Note: If you have projects without `go.mod` files that you want to add to the workspace, change into their project directory and run `go mod init`, then add the new module to your workspace with `go work use [path-to-module].`

## Workspace commands

Along with `go work init` and `go use`, Go 1.18 introduces the following commands for workspaces:

- `go work sync`: pushes the dependencies in the `go.work` file back into the `go.mod` files of each workspace module.
- `go work edit`: provides a command-line interface for editing `go.work`, for use primarily by tools or scripts.

Module-aware build commands and some `go mod` subcommands examine the `GOWORK` environment variable to determine if they are in a workspace context.

Workspace mode is enabled if the `GOWORK` variable names a path to a file that ends in `.work`. To determine which `go.work` file is being used, run `go env GOWORK`. The output is empty if the `go` command is not in workspace mode.

When workspace mode is enabled, the `go.work` file is parsed to determine the three parameters for workspace mode: A Go version, a list of directories, and a list of replacements.

Some commands to try in workspace mode (provided you already know what they do!):

```
go work init
go work sync
go work use
go list
go build
go test
go run
go vet
```

## Editor experience improvements

We’re particularly excited about the upgrades to Go’s language server [gopls](https://pkg.go.dev/golang.org/x/tools/gopls) and the [VSCode Go extension](https://marketplace.visualstudio.com/items?itemName=golang.go) that make working with multiple modules in an LSP-compatible editor a smooth and rewarding experience.

Find references, code completion, and go to definitions work across modules within the workspace. Version [0.8.1](https://github.com/golang/tools/releases/tag/gopls%2Fv0.8.1) of `gopls` introduces diagnostics, completion, formatting, and hover for `go.work` files. You can take advantage of these gopls features with any [LSP](https://microsoft.github.io/language-server-protocol/)-compatible editor.

#### Editor specific notes

- The latest [vscode-go release](https://github.com/golang/vscode-go/releases/tag/v0.32.0) allows quick access to your workspace’s `go.work` file via the Go status bar’s Quick Pick menu.

![Access the go.work file via the Go status bar&rsquo;s Quick Pick menu](https://user-images.githubusercontent.com/4999471/157268414-fba63843-5a14-44ba-be82-d42765568856.gif)

- [GoLand](https://www.jetbrains.com/go/) supports workspaces and has plans to add syntax highlighting and code completion for `go.work` files.

For more information on using `gopls` with different editors see the `gopls`[ documentation](https://pkg.go.dev/golang.org/x/tools/gopls#readme-editors).

## 下一步做什么？

- 下载和安装Go 1.18](https://go.dev/dl/).
- Try using [workspaces](https://go.dev/ref/mod#workspaces) with the [Go workspaces Tutorial](https://go.dev/doc/tutorial/workspaces).
- If you encounter any problems with workspaces, or want to suggest something, file an [issue](https://github.com/golang/go/issues/new/choose).
- Read the [workspace maintenance documentation](https://pkg.go.dev/cmd/go#hdr-Workspace_maintenance).
- Explore module commands for [working outside of a single module](https://go.dev/ref/mod#commands-outside) including `go work init`, `go work sync` and more.



## 后记



## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。https://www.zhihu.com/people/thucuhkwuji)



## References

* https://go.dev/blog/get-familiar-with-workspaces