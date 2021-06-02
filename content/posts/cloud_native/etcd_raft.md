---
title: "etcd_raft"
date: 2021-03-05T11:33:56+08:00
draft: false
isCJKLanguage: true

tags: ["cloud native"]

---


raft基本思想：少数服从多数。

leader 选举
+ 发起投票
  + 每个节点为follower， 每个节点均有election timeout到期时间，随机几百毫秒， 到期后身份变为candidate
  + candidate 发起投票
+ 投票过程
  + term 大
  + 日志最新
+ 当选后 发心跳包维护leader身份 


一次写请求：node.Propose --> MsgProp --> 加入log里，向follower广播； follower --> 收到 -> msg app resp
leader --> 收到多数resp --> commit log --> 向folower广播。

