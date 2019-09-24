# go
个人golang常用工具包
## proxy 设置

随着 go1.11 的发布，go 官方引入了 go module 来解决依赖管理问题，go module 被集成到原生的 go cmd 中，但是如果你的代码库在$GOPATH中，go1.11 的 module 功能是默认不会开启的，想要开启也非常简单, 通过一个环境变量即可开启go module：
```bash
export GO111MODULE=on

```

当我们使用go的时候，go默认会直接从代码库中去下载所需的相关依赖，GOPROXY 这个环境变量可以让我们控制自己从哪里去下载源代码，如果 GOPROXY 没有设置，go 会直接从代码库下载相关依赖代码。如果你像下面这样设置了这个环境变量，那么你就会通过 goproxy.io 下载所有的源代码。

```bash
export GOPROXY=https://goproxy.io
```
你可以通过置空这个环境变量来关闭，`export GOPROXY= `。

以前大家执行 go get golang.org/x/net net代码库会下载到本地GOPATH中，以后有任何项目引用到了 golang.org/x/net 都不会再去下载这个代码库，因为本地GOPATH已经有了，哪怕版本不对，golang也会引用。但是随着 module 概念引入go语言，每个引入的 module 拥有了 version。随着代码库的不断更新迭代，大家即使是对同一个代码库的引用也可能用了不同的tag 或者 commit hash，基于这个现状，go1.11 的 module 会比以前更频繁的下载源代码。但是基于中国有中国特色的互联网，我们有时候很难get到我们需要的依赖源代码，进而导致项目编译失败，CI失败。于是，我们需要一个proxy。

## how to create
```bash
echo "# golib4db" >> README.md
git init
git add README.md
git commit -m "first commit"
git remote add origin https://github.com/bigdot123456/golib4db.git
git push -u origin master

```
## how to use

*just open go.mod*

### include which lib ?

1. redis
2. db linker
3. crypto
4. error dispose
5. log
