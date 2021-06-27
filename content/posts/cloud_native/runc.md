---
title: "runc"
date: 2021-06-26T18:34:48+08:00
draft: false
isCJKLanguage: true
tags: ["cloud native"]
series: [""]
categories: ["技术"]
---

runc 是根据OCI（Open Container Initiative）开放容器标准开发的容器运行工具。

runc 在libcontainer上进行各种命令的封装。

## 1. 项目结构

+ runc/main.go 文件是命令的入口文件，其余子命令模块也存在于该层目录。

如 create.go，start.go, run.go exec.go 等；整个命令模块的组装使用了 github.com/urfave/cli 这个库

+ runc/libcontainer 目录是主要存放每个子命令模块的工作主体逻辑；

## 2. 命令

### 2.1 create, init, start

+ create：解析配置和创建子进程的消息通道
+ init：根据容器配置启动容器运行环境，阻塞写在execfifo队列，等待start的读，随后execve执行用户程序
+ start：为了通知init 进程启动容器

<img src="/cloud_native/runc_run.svg" width = "300" /><br>


create命令主要2个步骤：

1. spec, err := setupSpec(context)

解析配置

2. status, err := startContainer(context, spec, CT_ACT_CREATE, nil)

startContainer方法在create, run restore命令中共用，只是入参action不同。

根据配置生成runc init容器进程

```golang
func (l *linuxStandardInit) Init() error {
    ......
    name, err := exec.LookPath(l.config.Args[0])
    // fd是execfifo, 写阻塞，等start执行读fifo后接着执行
    if _, err := unix.Write(fd, []byte("0")); err != nil {
		return newSystemErrorWithCause(err, "write 0 exec fifo")
	}
    // init进程执行exec后，替换成用户程序
    if err := unix.Exec(name, l.config.Args[0:], os.Environ()); err != nil {
		return newSystemErrorWithCause(err, "exec user process")
	}
}
```

### 2.2 run

runc run containerID, 少了execfifo的阻塞，过程类似create, start。

pause/resume利用了 cgroup 的 子系统 freezer 来实现进程的挂起

---

## 3. namespace

nsexec.c中从”_LIBCONTAINER_INITPIPE”环境变量中拿到pipe，从pipe中并读取bootstrapData。
先执行clone()，参数有CLONE_PARENT及命名空间参数，使用子进程和父进程是兄弟关系，并拥有自己的命名空间。
然后调用setns()加入存在的namespace

## 4. cgroup

cgroup对进程资源的限制，在libcontainer中是往子系统下的资源控制文件里写入配置值。


## 5. ebpf 

限制在容器中访问设备以及访问权限，不是往devices子系统里写配置值，而是采用ebpf限制访问。 

一段字节码加载到内核中执行。

github.com/cilium/ebpf


