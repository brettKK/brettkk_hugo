---
title: "kubernetes"
date: 2021-03-03T18:34:48+08:00
draft: false
isCJKLanguage: true
tags: ["cloud native"]
series: [""]
categories: ["技术"]
---

### 容器技术历史

1979年 chroot 设置进程访问的根目录

2006年 google 给kernel提供Control Groups 

2008年 LXC 第一个容器管理方案

2015年 google 提供libcontainer

2013年 容器管理产品docker

---

cgroup： 程序运行时对资源的调度触发相应的钩子以达到资源追踪和限制的目的

Open Container Initiative（OCI）成立后， libcontainer 封装在runC包里。

---

+ pod
  + 管理关系紧密的容器进程，共享ipc network
  + 容器不只containerd，k8s需要基于container做一步抽象container runtime interfz
+ service
  + 一个svc是一个微服务，也是一组prod，通过label robin实现IP负载均衡
+ volume