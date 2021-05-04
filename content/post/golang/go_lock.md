---
title: "go lock"
date: 2021-05-03T18:34:48+08:00
draft: false
isCJKLanguage: true
tags: ["golang"]
series: [""]
categories: ["技术"]
---

locks 锁

### 源码包的结构

+ sync包
  + sync.Mutex
  + sync.RWMutex
  + sync.Cond
  + sync.WaitGroup
  + sync.Once
  + sync.Pool 涉及的内容很多
    + cache line padding
    + no copy
    + lock-free
    + victim cache
    + gc 细节

### 写一个锁

cas是基于硬件lock and swap 实现，早期386会lock总线，实现swap时无其他竞争，后期lock总线细粒度到cpu的cache里，通过MESI一致性保证swap时的正确性。

```
for {
	复制旧数据
	基于旧数据构造新数据
	if CompareAndSwap(内存地址，旧数据，新数据) {
		break
	}
}
```

```
var flag int
func Reader() {
    for {
        if atomic.CompareAndSwap(&flag, 0, 1) {
            // do other thing
            atomic.Store(&flag, 0)
            return
        }
        // cas failed, do again
        // simple spinlock
    }
}

```



+ futex
+ spin vs sleep
    + cas spin几次，没拿到锁则进入futex syscall后sleep。
    + sync.Mutex的实现方式， cas spin几次，没拿到锁则进入enqueue waitqueue后sleep。

+ sync.Mutex
   + 实现了公平锁，公平的意思是排队等待锁
   + 2种模式，正常模式 和 饥饿模式
   + 第一次cas失败后，进入lockslow方法（为了主方法Lock()能inline进去上层代码里），spin几次失败后，进去队列等待被唤醒
