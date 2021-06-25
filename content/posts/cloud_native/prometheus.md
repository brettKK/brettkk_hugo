---
title: "prometheus"
date: 2021-06-24T18:34:48+08:00
draft: false
isCJKLanguage: true
tags: ["cloud native"]
series: [""]
categories: ["技术"]
---

+ 时序数据库
    + TSDB
    + influxDB

## 入口main方法

### 管理协程的生命周期

使用github.com/oklog/run。

<img src="/cloud_native/prom_rungroup.svg" width = "600" /><br>

场景：管理多个子协程，通过main goroutine的errors channel返回子协程的execute error， 中断所有的子协程。