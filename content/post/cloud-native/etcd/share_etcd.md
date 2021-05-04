+ Etcd是如何实现一致性的
+ Etcd的存储是如何实现的
+ watch机制，key过期机制

+ etcd 简介
+ raft协议原理 （保证写操作的改动会可靠的同步到其他server节点）
  + leader选举
  + 日志复制
  + 安全性
+ 操作实践
+ 源码分析
  + 客户端源码（键值，集群配置，watcher机制）
  + 服务端源码（启动http服务，响应客户端请求）
  + 集群间通讯
  + 数据存储

---
客户端源码

1 负载均衡器

2 客户端api
  返回均含response header {cluster_id, member_id, revision, raft_term}

3 错误处理器


---
从paxos到zookeeper分布式一致性原理与实践
+ cap 一致性 可用性(有效时间内，返回结果) 分区容错性
+ base basically available, soft state, eventually consistent
+ 一致性协议
  + 2PC 两阶段提交
  + 3PC 三阶段提交
  + paxos
  
--- 

in company
+ do one plan， 方法论，流程规范论
+ 8 - 10 总结 学习技术, 源码阅读, ppt介绍技术 技术分享～～ 的流程模式
+ 10 -12 熟悉业务背景，数据流，逻辑过程  完成需求
+ 14 - 16 完成需求
+ 17 - 18.5
+ 19.5 - 22.0 have not reach

---
+ learn tech and write ppt (etcd has ppt)
+ language
+ ml cs231n
