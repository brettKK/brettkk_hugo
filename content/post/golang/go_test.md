---
title: "go test todo"
date: 2021-05-03T18:34:48+08:00
draft: false
isCJKLanguage: true
tags: ["golang"]
series: [""]
categories: ["技术"]
---


+ goland need money, so vs code 
+ Configure Visual Studio Code for Go
    + extensions: go
    + view -> command palette... -> goinstall



+ pprof (CPU profiles, Heap profiles, block profile, traces)
    + kite, ginex 框架开启了prof功能， 定时任务，消费者等worker需要 import _ net/http/pprof手动开启
    + pprof 性能 ( CPU profile, Mem profile, Block Profiling, Goroutine Profiling )
        + go tool pprof 
            + go tool pprof http://localhost:8080/debug/pprof/profile
        + http://127.0.0.1:8080/debug/pprof/heap?debug=1 
        + 终端连接到 go tool pprof -inuse_space http://127.0.0.1:8080/debug/pprof/heap
            + 进去pprof终端， top 10 或者web生成调用关系图svg
            + list
            + disasm
        + svg go tool pprof -alloc_space -cum -svg http://127.0.0.1:8080/debug/pprof/heap > heap.svg
        + go-torch -alloc_space http://127.0.0.1:8080/debug/pprof/heap --colors=mem
        + go-torch http://localhost:8080/debug/pprof/profile
    + runtime/pprof
    + net/http/pprof
+ wrk 压力测试 github.com/juju/ratelimit 

+ golang 测试框架 stub/mock
  + GoConvey  https://github.com/smartystreets/goconvey
    + cd <your project> and $GOPATH/bin/goconvey, in browser http://localhost:8080
    + Convey, so, 不是很习惯，找不到定义的地方
  + testify  https://github.com/stretchr/testify
  + GoStub 
  + GoMock  只能mock interface里的函数
  + Monkey


+ delve 调试工具 

+ 优化go程序 运行速度的方向 只对关键的路径进行优化 (优化运行速度，开会速度会慢 hah)
    + channel（mutex锁保证并发安全） ==》 ringbuffer (CAS 保证并发安全)
    + 避免堆分配可以成为优化的主要方向. 若频繁使用堆，可以sync.Pool 复用对象
    + cpu cache line(64字节) False Sharing。为了保证cache的一致性，对内存的一个小小的写入都会让cache line被淘汰。对相邻地址的读操作就无法命中对应的cache line。 RingBuffer struct padding。  type RingBuffer struct {_padding0  [8]uint64, queue          uint64, 	_padding1      [8]uint64, 	dequeue     uint64}