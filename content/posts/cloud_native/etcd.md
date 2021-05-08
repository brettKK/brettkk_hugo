---
title: "etcd"
date: 2021-04-05T11:33:56+08:00
draft: false
isCJKLanguage: true

tags: ["cloud native"]

---

+ etcd架构分为4部分
  + http server
  + store
  + raft
  + wal

<!-- ![image](/cloud_native/etcd_arch.jpg) -->
<img src="/cloud_native/etcd_arch.jpg" width = "400" /><br>

+ etcd使用场景  
  + API：Etcd提供HTTP+JSON, gRPC接口，跨平台跨语言; 支持https
  + 共享配置
  + 服务发现
  + 选主
  + 分布式队列
  + 分布式锁

+ play.etcd.io, dash.etcd.io

+ 数据存储
  + 内存存储，顺序记录数据的变更记录，也会有索引和建堆等方便查询
  + 持久化到硬盘， wal进行记录的存储(wal目录和snapshot)
+ etcd wal日志格式，日志命令规则 (tools/etcd-dump-logs)
  + WAL文件名以$seq-$index.wal的格式存储, 日志切分seq加1，新操作index加1
  + entry
    + type 0-normal, 1-confchange
    + term 主节点的任期 （从一次竞选开始到下一次竞选开始）
    + index 序号，严格有序递增 snapshot的存储命名以$term-$index.wal格式进行命名存储
    + data


+ etcd (gopher redis + zookeeper) 一致性的kv存储系统
  + raft，复制日志的一致性算法 (选举，日志复制，安全性)
    + 选举，理解raft的时钟周期和超时机制
    + 日志复制，理解日志的同步机制
    + 主节点处理所有的写操作，通过raft协议可靠的同步到其他节点