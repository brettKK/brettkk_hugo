---
title: "kubernetes"
date: 2021-03-03T18:34:48+08:00
draft: false
isCJKLanguage: true
tags: ["cloud native"]
series: [""]
categories: ["技术"]
---


+ pod
  + 管理关系紧密的容器进程，共享ipc network
  + 容器不只containerd，k8s需要基于container做一步抽象container runtime interfz
+ service
  + 一个svc是一个微服务，也是一组prod，通过label robin实现IP负载均衡
+ volume