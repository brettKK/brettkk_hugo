---
title: "k8s 学习"
date: 2021-06-27T11:33:56+08:00
draft: false
isCJKLanguage: true

tags: ["cloud native"]

---


## k8s架构

<img src="/cloud_native/k8s_architecture.png" width = "600" /><br>



## 请求访问pod的流程


hostNetwork： 占用特定宿主机上的特定端口

hostPort： 将容器的端口与node上的端口做映射
NodePort
LoadBalancer
Ingress

## 网络