---
title: "redis 小结"
date: 2021-05-05T11:33:56+08:00
draft: false
isCJKLanguage: true
tags: ["存储"]
series: [""]
categories: ["技术"]
---

+ 支持的value类型： 字符串，hash，列表，集合，有序集合
+ redis未直接使用上述数据结构， 而是redisObject结构进行封装。
    + 键对象， 均为字符串对象
    + 值对象， 为1字符串对象，2列表对象(ziplist, linkedlist)，3哈希对象(ziplist, hashtable)，4集合对象（intset, hashtable），5有序集合对象(ziplist, skiplist)
    + redisObject结构体： type, encoding, prt, refcount, lru...

+ 简单动态数组
    + sds.h sds.c
    + 数组本身长度， 空余长度，防止strcat 这种缓冲区溢出
    + 兼容c字符串库函数的原因: SDS返回char*指针与C 操作字符串时是一致
  
+ 压缩列表 ziplist
    + 组成 <zlbytes, zltail, zllen, entry1,entry2,...entryN, zlend>
    + 整体长度-字节 == 指向tail的偏移量 == entry的个数 == entry... == 标记结尾0xff 一字节
    + entry的结构<previous_entry_length, encoding(content的类型), content>, previous_entry_length 支持从尾到头
    + previous_entry_length带来的连锁更新问题

+ 过期键删除
    + 定时删除
        + 创建定时器添加到时间事件的链表中
    + 懒性删除策略(db.c/expireIfNeeded) get key -> expireIfNeeded -> 删除key 并返回nil
    + 定期删除策略 (redis.c/activeExpireCycle)
    + save 生成rdb文件， 不会保存过期键

+ RDB持久化
    + save (阻塞redis服务进程，直到RDB文件生成完毕) 
    + bgsave 创建子进程来创建全局快照RDB文件


+ AOF持久化
    + 配置appendfsync
        + always 每次写都fsync
        + everysec 每秒fsync
        + no  由os定
    + AOF 重写（减少AOF文件的大小）， bgrewriteaof创建子进程去执行
    + 重写： 创建新AOF文件，通过读当前db的键的value， 然后用一条命令去记录key value 信息
    + 子进程在AOF重写过程中， 发生的write操作，主服务进程会写到AOF重写缓冲区，等子进程结束后，父进程收到信号，将AOF重写缓冲区的内容append到AOF文件中，完成AOF文件的后台重写。
    + AOF 压缩原理：..

+ 事件
    + 文件事件 ae_select.c , ae_epoll.c, ae_kqueue.c reactor模式
    + 时间事件 (定时事件， 周期性事件)
        + 更新db的统计信息
        + 清理过期key
        + 关闭和清理无效的客户端
        + 尝试AOF RDB
        + 主从同步

+ redis cluster 模式
    + 数据分布 节点取余  一致性哈希
    + 虚拟槽分区 16384个槽的分配
    + moved错误（请求a到机器上，但机器存放该key对应的槽，所以返回集群中正确的地址）

+ 事务
    + multi, exec, watch
    + multi之后 的命令入队列， exec执行队列命令。
    + multi之后 redisclient的flags事务标示打开， 请求的命令入队列（入队失败，不会执行事务）， exec执行队列命令（执行失败，事务继续执行）。
    + watch的key被修改后，redisClient的REDIS_DIRTY_CAS会打开， exec时会取消执行
    + redis单线程处理事件，所以事务执行是串行的，所以满足隔离性。

+ redis中SIGUSR1信号关闭子进程的aof,rdb任务

缓存穿透解决方案
+ 缓存空值，一次
+ 布隆过滤器

+ exactly once 有且仅有一次，（与外部存储系统协作）
  + 生产者发消息时附带pid+messageid， broker检查messageid是否比我的messageid大于1， 是则接受；不是拒接
  + 消费者更新数据和更新offet需要事务支持，一起成功或者一起失败，保证正好消费一次