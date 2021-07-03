---
title: "prometheus"
date: 2021-06-24T18:34:48+08:00
draft: false
isCJKLanguage: true
tags: ["cloud native"]
series: [""]
categories: ["技术"]
---

在解决内存与磁盘的读写模式、性能的不匹配问题。有以下两点：

在内存中的head block进行写，使用wal保证内存数据的持久性。

后台协程定期compact到磁盘的block文件夹，使用index文件来提高查询效率。


## 1. 入口main方法

### 多协程的共进退

使用github.com/oklog/run。

<img src="/cloud_native/prom_rungroup.svg" width = "600" /><br>

场景：管理多个子协程，通过main goroutine的errors channel返回子协程的execute error， 中断所有的子协程。


## 2. tsdb

+ chunks
    + sample, raw label set, timestamp & value tuple
    + sequential series of encoded samples
+ chunk head 内存存储
+ wal
+ block 
    + mini database: index, meta, tombstone, chunks
    + two key processes: compaction, truncation
+ index
    + TOC段， 存放symbol段，series段，posting倒排索引段在index文件的offset
    + 倒排索引，label1, value1 -> serie id list
    + series段,  label sets -> chunks
    + 查询过程； 根据多个label， value查询posting倒排索引，取serie id的交集，再根据serie id查找series段查找到chunk信息。