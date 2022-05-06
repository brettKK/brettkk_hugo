---
title: "bug_fix"
date: 2021-05-04T19:10:43+08:00
draft: false
isCJKLanguage: true
tags: ["golang"]
series: [""]
categories: ["技术"]
---

一些bug
+ init() 函数滥用
+ 函数功能是查询类的，能不用指针传参数就不用，避免函数内部对对象做修改
+ 代码里不要吞掉任何一个错误
+ first path segment in URL cannot contain colon 原因：字符串http前面多了空格，post请求在本地就会失败
+ err:crypto/cipher: input not full blocks
    + AES加密后的数据存入mysql时，由于字段有长度限制， 加密后的内容被截断了，所以AES解密时爆出这个panic
+ 问题排查 http
    + 499对应的是 “client has closed connection”。这很有可能是因为服务器端处理的时间过长，客户端主动关闭
    + 字段总结 remote_addr, x_forwarded_for
+ kafka: the provided member is not known in the current generation kafka
    + rebalance
+ signal.Notify(signals) （Notify使用时需要指定接收信号的类型）， 这种写法 <-signals可以接收os发的64种信号, 其中 broken pipe signal 导致进程主进程结束
    + 本质是 IOError 错误, 读写文件IO和网络Socket IO
+ 管道存在上游发送数据的进程，下游读取数据的进程，在下游不再需要读取上游数据的时候，就会发送 SIGPIPE 信号给上游进程
    + Broken Pipe 是可以忽略
+ .ignore文件中配置dubug，导致忽略了vendor里的dubug包，导致本地编译通过，但是远程scm编译失败
+ brew upgrade go && should modify .zshrc
+ ss -l看到主机上的端口在使用， lsof -i:port 没有发现使用这个端口的进程
+ Unix domain socket 同一台操作系统上的两个或多个进程进行数据通信 跨进程通信 
+ 使用 循环迭代器变量的引用， 大概率全部遍历的都是迭代的最后一个元素
+ go-zero, bloom filter. core/bloom/bloom.go  


---
unsolved: 阅读go源码时，执行单测出现： use of internal package internal/byteal not allowed