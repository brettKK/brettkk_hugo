---
title: "lsm 小结"
date: 2021-03-03T18:34:48+08:00
draft: false
isCJKLanguage: true
tags: ["存储"]
series: [""]
categories: ["技术"]
---
## lsm 基本想法

使用场景， 写多读少，例如存调用链的服务。 

应用：levelDB, HBase, Cassandra, RocksDB

+ 为什么lsm
  + 分层，有序，面向磁盘的数据结构
  + 批量写，顺序写对硬件驱动友好
  + 顺序操作高吞吐

## 写入

写数据时， 写入内存有序树结构memtable, 并更新布隆过滤器。

当memtable累积到阈值，一次性写入segment内部有序文件到磁盘。

后台有专门的定期执行compact操作的线程合并segment文件。


## 读取

查询布隆过滤器，若不存在则一定是不存在； 若存在，按照从新到老的顺序依次查询每个segment。

## 更新
## 删除

删除时是覆盖写value为特殊标识位，而不是真执行查找并删除。