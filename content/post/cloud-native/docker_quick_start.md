+ docker 三个核心概念
    + 镜像image
    + 容器container
    + 仓库repository

+ advantages
  + same evironment
  + sandbox projects
  + it just works

+ container
  + run image
    + OS
    + software
    + application code
+ image
  + defined by a docker file
    + configure operating system
    + install the software
    + etc

+ docker for mac
    + Docker Toolbox (virtualbox, docker客户端， docker compose, kitematic gui, docker machine)  /usr/local/bin 
    + github.com/docker/toolbox/releases
+ hub docker com

+ console demand
  + -p 映射容器中的端口到主机的端口 -P
  + -v 挂载一个磁盘卷
  + -t pseudo-TTY
  + docker build -t hello-world .
  + docker run -p 80:80 hello-world
  + docker run -p 80:80 -v ~/src:/var/www/html/ hello-world
  + docker-compose up
  + docker-compose up -d
  + 在docker中创建一个bash shell
    + docker run --name test -it debian

+ 容器
  + cgroup 资源控制
  + namespace 访问控制
    + /proc/[pid]/ns目录下有该进程对应的namespace的inode
    + /proc/[pid]/mounts 文件保存了进程所在 namespace 所有已经 mount 的文件系统
    + namespace是进程的属性
  + rootfs 文件系统隔离 文件系统的挂载点 可访问的文件
    + UnionFS:AUFS, btrfs, vfs, DeviceMapper
  + 容器引擎 生命周期控制
  + container format
    + namespaces
    + control groups
    + UnionFS

+ OCI(open container initiative)
    + image spec 
        + 文件系统 
        + config文件 保存了文件系统的层级信息
        + manifest文件 镜像的 config 文件索引
        + index文件 
    + runtime spec

+ 虚拟机是如何实现的
+ docker是如何实现的
  + namespace
  + cgroup
  + union file system
+ 容器是如何被创建出来的

--- 
kubernetes 
    + Minikube 运行本地单节点的k8s集群
        + minikube start