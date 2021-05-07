---
title: "go_toolchain"
date: 2021-05-01T18:54:41+08:00
draft: false
isCJKLanguage: true
tags: ["golang"]
series: [""]
categories: ["技术"]
---

### an overview of go's tooling

+ table of content 
+ installing tooling
    + $ GO111MODULE=on go get golang.org/x/tools/cmd/stresss
    + go get -u, 强制使用网络去更新包和它的依赖包, 自动完成编译和安装 (git clone + go install)
+ viewing environment information
    + go env
    + go env GOPATH GOOS GOARCH
    + go help environment (go help get)
+ development
    + GOOS=windows go build
        + go build -gcflags="-m -l" // 看变量分配在stack or heap
    + go run ./cmd/foo
    + go run, go test, go build 可以静默添加外面依赖
    + go get github.com/foo/bar@v1.2.3 也可以指定特别的版本号
        + go get -v 显示操作流程的日志及信息，方便检查错误
        + go get -u 重新下载依赖包
        + go get -d 只下载，不安装
        + go get -insecure  允许使用不安全的 HTTP 方式进行下载操作
    + go list -m all
    + go mod why -m golang.org/x/sys
    + go mod graph
    + go clean -modcache (cache: GOPATH/pkg/mod)
    + gofmt -d -w -r 'Add -> ADD' .
    + go doc -src strings.Replace (view go documentation) , go doc fmt Printf
+ testing
    + go test ./foo/bar
    + go clean -testcache
    + go test -v -run=^TestFooBar$ .
    + go test -cover ./...
    + go test -run=^TestFooBar$ -count=500 . (stress testing)
    + go test -covermode=count -coverprofile=/tmp/profile.out ./...
    + go tool cover -html=/tmp/profile.out
+ pre_commit checks
    + gofmt -w -s -d .
    + go vet foo.go
    + golint foo.go
    + go mod tidy, go mod verify
+ build and deployment
    + go clean -cache 
    + go install -> $GOPATH/bin/demo
    + go list -f '{{ .Name}}'
    + go tool dist list (os/architecture)
    + GOOS=linux GOARCH=amd64 go build -o=/tmp/linux_amd64/foo .
    + https://github.com/golang/go/wiki/WebAssembly
+ diagnosing problems and making optimizations
    + go tool pprof --nodefraction=0.1  -http=:9999 /tmp/cpuprofile.out
    + go tool trace
        + goroutine的时间阶段： Execution Time， Network Wait Time， Sync Block Time， Blocking Syscall Time， Scheduler Wait Time， GC Sweeping， GC Pause。