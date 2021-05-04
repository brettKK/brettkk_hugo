---
title: "golang内存模型"
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


  + Goroutinue Creation  happens before
  + Goroutinue destruction no happens before
  + Channel communication
    + the main method of Synchronization between goroutines
  + Locks
    + sync package implements two lock: sync.Mutex and sync.RWMutex
  + Once  单例dao对象


参考： golang.org/ref/mem
