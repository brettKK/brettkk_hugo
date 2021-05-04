+ pod
  + 管理关系紧密的容器进程，共享ipc network
  + 容器不只containerd，k8s需要基于container做一步抽象
+ service
  + 一个svc是一个微服务，也是一组prod，通过label robin实现IP负载均衡
+ volume