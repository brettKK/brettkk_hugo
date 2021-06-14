---
title: "go内存模型"
date: 2021-05-03T18:34:48+08:00
draft: false
isCJKLanguage: true
tags: ["golang"]
series: [""]
categories: ["技术"]
---

go的内存模型规定： 在哪些情况下，一个golang中读取的变量时可以观测到另一个golang对同一变量的新写的值。

+ init 方法的执行顺序

![image](/golang/init_func_order.png)

+ goroutine创建
  + 创建一个golang happens before 新golang的执行

```golang
var a string
func f() {
  print(a)
}
func hello() {
  a = "hello, world"
  go f() // 等到f执行时打印hello world
}
```

+ goroutine销毁 没有happen before任何程序里的代码
```golang
var a string
func hello() {
  go func() {a = "mike"}
  print(a) // 不确定的值
}
```

+ channel通信
  + 写channel的动作 happen before 从channel的读
  + 关闭channel的动作 happen before 从channel中接收零值
  + 从非缓冲channel接收的动作 happen before 写channel

+ Locks 
  + sync.Mutex and sync.RWMutex的unlock happen before lock

+ Once  单例对象
  + 第一次once.Do(f) happen before 后面无数次调用Do(f)
  

参考： golang.org/ref/mem
