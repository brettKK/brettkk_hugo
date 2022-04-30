---
title: "go test"
date: 2021-05-05T18:34:48+08:00
draft: false
isCJKLanguage: true
tags: ["golang"]
series: [""]
categories: ["技术"]
---

## 工具
+ Guru 导航go 代码的编辑器集成工具
    + 变量，函数的声明地点
    + 变量，函数的所有引用地点
    + 实现此接口的所有具体类型
+ golang.org/x/tools/cmd/guru
    + http://golang.org/s/using-guru
## 单测

单测mock框架，测试框架 stub/mock
+ GoConvey  https://github.com/smartystreets/goconvey
    + Convey, so, 不是很习惯
+ testify  https://github.com/stretchr/testify
  + GoStub 
  + GoMock
+ https://bou.ke/
  + copy 覆盖函数的前12个字节的汇编代码，植入mock函数的地址， 在运行时实现函数的mock和unmock
  + github.com/bouk/monkey

### gomock

install: go get github.com/golang/mock/gomock, github.com/golang/mock/mockgen.



### go monkey


```
package main
func from() int {return 1}
func to() int (return 2)
func main() {
  // 在执行from之前把from函数内的机器码 替换为一条跳转指令，让cpu跳转到to函数的机器码上执行。
  patch(from, to)
  fmt.Println(from()) // should print 2
}

```
taoshu.in/go/monkey
+ 找到from, to的内存地址
+ 修改from函数的机器码，构造跳转指令，跳到to的内存地址上

---
## 调试

+ delve 调试工具 
+ gdb单步调试go程序的执行过程    Delve 更好
  
---
## 性能分析

利用runtime/pprof， net/http/pprof采集运行时数据

+ pprof (CPU profiles, Heap profiles, block profile, traces)
    + kite, ginex 框架开启了prof功能
    + 定时任务，消费者等worker需要 import _ net/http/pprof手动开启
    + 性能pprof
        + go tool pprof http://localhost:8080/debug/pprof/profile
        + http://127.0.0.1:8080/debug/pprof/heap?debug=1 
        + 终端连接到 go tool pprof -inuse_space http://127.0.0.1:8080/debug/pprof/heap
        + svg go tool pprof -alloc_space -cum -svg http://127.0.0.1:8080/debug/pprof/heap > heap.svg
+ wrk 压力测试 github.com/juju/ratelimit 

+ 优化go程序 运行速度的方向 只对关键的路径进行优化 (优化运行速度，开会速度会慢 hah)

1. 避免堆分配可以成为优化的主要方向。 若频繁使用堆，可以sync.Pool 复用对象
2. cpu cache line(64字节)，结构体填充，避免False Sharing。
3. 为了保证cache的一致性，对内存的一个小小的写入都会让cache line被淘汰。
4. 对相邻地址的读操作就无法命中对应的cache line。 