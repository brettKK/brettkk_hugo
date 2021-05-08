---
title: "go test"
date: 2021-05-03T18:34:48+08:00
draft: false
isCJKLanguage: true
tags: ["golang"]
series: [""]
categories: ["技术"]
---

### 单测

+ golang 测试框架 stub/mock
  + GoConvey  https://github.com/smartystreets/goconvey
    + Convey, so, 不是很习惯
  + testify  https://github.com/stretchr/testify
  + GoStub 
  + GoMock
  + github.com/bouk.monkey
  + https://bou.ke/
    + copy 覆盖函数的前12个字节的汇编代码，植入mock函数的地址， 可以在运行时实现函数的mock和unmock

---
### 调试

+ delve 调试工具 
+ gdb单步调试go程序的执行过程    Delve 更好
+ Guru 导航go 代码的编辑器集成工具
    + 变量，函数的声明地点
    + 变量，函数的所有引用地点
    + 实现此接口的所有具体类型
+ golang.org/x/tools/cmd/guru
    + http://golang.org/s/using-guru


---
### 性能分析

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
    + runtime/pprof
    + net/http/pprof
+ wrk 压力测试 github.com/juju/ratelimit 

+ 优化go程序 运行速度的方向 只对关键的路径进行优化 (优化运行速度，开会速度会慢 hah)
    + channel（mutex锁保证并发安全） ==》 ringbuffer (CAS 保证并发安全)
    + 避免堆分配可以成为优化的主要方向. 若频繁使用堆，可以sync.Pool 复用对象
    + cpu cache line(64字节) False Sharing。
    + 为了保证cache的一致性，对内存的一个小小的写入都会让cache line被淘汰。
    + 对相邻地址的读操作就无法命中对应的cache line。 