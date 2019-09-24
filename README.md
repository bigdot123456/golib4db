# go
个人golang常用工具包
------------
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
------------

*just open go.mod*
------------

### include which lib ?

1. redis
2. db linker
3. crypto
4. error dispose
5. log


## FAQ

* I/O Timeout Errors 
The golang.org domain may be blocked from some countries. go get usually produces an error like the following when this happens:

```sh
$ go get -u google.golang.org/grpc
package google.golang.org/grpc: unrecognized import path "google.golang.org/grpc" (https fetch: Get https://google.golang.org/grpc?go-get=1: dial tcp 216.239.37.1:443: i/o timeout)
To build Go code, there are several options:
```
### solving method
1. Set up a VPN and access google.golang.org through that.

2. Without Go module support: git clone the repo manually:

```bash
git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc

```
You will need to do the same for all of grpc's dependencies in golang.org, e.g. golang.org/x/net.

3. With Go module support: it is possible to use the replace feature of go mod to create aliases for golang.org packages. In your project's directory:
```bash

go mod edit -replace=google.golang.org/grpc=github.com/grpc/grpc-go@latest
go mod tidy
go mod vendor
go build -mod=vendor

```

Again, this will need to be done for all transitive dependencies hosted on golang.org as well. Please refer to this issue in the golang repo regarding this concern.

### Compiling error, undefined: grpc.SupportPackageIsVersion

Please update proto package, gRPC package and rebuild the proto files:
```bash

go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/grpc
protoc --go_out=plugins=grpc:. *.proto

```

### How to turn on logging
The default logger is controlled by the environment variables. Turn everything on by setting:
```bash
GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info

```

The RPC failed with error "code = Unavailable desc = transport is closing"
This error means the connection the RPC is using was closed, and there are many possible reasons, including:

1. mis-configured transport credentials, connection failed on handshaking
1. bytes disrupted, possibly by a proxy in between
1. server shutdown 

It can be tricky to debug this because the error happens on the client side but the root cause of the connection being closed is on the server side. Turn on logging on both client and server, and see if there are any transport errors.