---
title: "kubernetes"
date: 2021-03-03T18:34:48+08:00
draft: false
isCJKLanguage: true
tags: ["cloud native"]
series: [""]
categories: ["技术"]
---


### service mesh

微服务内部有分布式环境下的通用功能：熔断策略、负载均衡、服务发现、认证和授权、quota限制、trace和监控等等 

这种微服务存在的问题： 
+ 虽然有现成的框架做了这些事情，但在也需要业务去配置和管理框架
+ 这些框架lib或者sdk缺业务语言时，很难融入或者需要开发对应语言的包来融入微服务内部
+ 框架升级，对业务不透明，需要业务自行升级

sidecar 单边模式代理模式。 同一个微服务的代理进程service mesh互通，与其他微服务的service mesh也通信，所以单看代理进程像网一样，所以叫service mesh。


### serviceless

主要包含两个领域的技术，BaaS（Backend as a Service）和FaaS（Function as a Service）

简化后端开发流程...

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