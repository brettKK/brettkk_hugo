---
title: "linux-base"
date: 2021-05-05T11:33:56+08:00
draft: false
isCJKLanguage: true
markup: mmark
tags: ["linux"]
series: [""]
categories: ["技术"]
---


### Linux 版本

uname -a 查看系统的版本
+ debian/ubuntu
+ redhat/fedora/centos

各种源码协议
+ 必须开源： LGPL Mozilla  GPL（不允许将 GPL 代码用于商业产品）
+ 可以闭源： BSD MIT Apache（宽松）

+ ls = list
	+ ls -a (display all files)
	+ ls -l (more infomation)
	+ ls -lh (file size k)

### 磁盘相关的命令
+ du 显示某个目录或者文件占用的磁盘空间
	+ 查看文件大小 与ls 的区别
+ df 显示整个磁盘的使用情况

+ find
	+ 以文件名查找 find . -name "xx.txt"
	+ 以文件从属关系查找 find . -user brett
	+ 以文件类型查找 find . -type f(普通文件，d目录文件，p管道文件)
	+ 以文件大小查找 find . -size -1M(小于1M的文件，+1M大于1M的文件)
	+ 以时间查找
		+ -atime 文件访问时间，天为单位
		+ -mtime 文件数据修改时间
		+ -ctime 文件元数据修改时间
	+ 指定查找深度	-maxdepth, -mindepth
	+ find与xargs结合
		+ find . -name "a.txt" | grep 'key'
		+ find . -name "a.txt" | xargs grep 'key'
		+ find . -name "a.txt" | xargs rm
	+ find -exec 与其他命令结合使用
		+ find . -name "a.txt" -exec rm {} \;


管道与xargs
+ 管道：将前面的标准输出作为后面的标准输入 echo "--help" | cat
+ xargs：将前面的标准输入作为后面命令的参数 echo "--help" | xargs cat

+ 文件特殊权限
	+ suid
	+ sgid
	+ sticky

+ 用户权限-文件
	+ chown = change owner
	+ chgrp = change group
	+ chmod = change mode

---

### 有显示功能的命令
+ cat
+ cut 按列分文件
+ head -n 5 a.txt
+ tail
	+ tail -n 10 a.txt
	+ tail -f xx.log | grep hello
+ 利用head和tail查看文件中3000-5000的行
	+ cat a.txt | head -n 5000 | tail -n 2000
	+ cat a.txt | tail -n + 3000 | head -n 2000
+ sed也可以查看文件中特定的行
	+ sed -n "3000,5000p" a.txt
+ more , less(更新一点)，功能上没看到区别
+ grep -i 匹配时忽略大小写, --color
+ sort
	+ -k  (按某一个列排序)
	+ -n （按数值降序排序， 不加-n可以出现10小于2的情况）
	+ -nr （按数字逆序排序）
	+ 文件A， B 差集（A - B） sort a.csv b.csv|uniq -u > c.csv
+ uniq = unique
	+ uniq -c  (相同行爱在一起 所以前面是sort命令)
+ wc = word count
	+ wc -c
	+ wc -l 	
+ awk， sed功能有些复杂，可以google search...

---

### 网络常用工具命令（文件夹network）
+ ss = socket statistics
  + 用来显示处于活动状态的套接字信息
  + 快，利用到了TCP协议栈中tcp_diag
  + ss -t 显示tcp连接
+ nc 网络发包工具
+ netstat
	+ netstat -an
+ tcpdump

---

### 进程相关的命令
+ ps
	+ -e 显示所有进程
	+ -f 全格式
	+ -l 长格式
	+ -L ps -L pid , 查特定进程pid的线程
	+ -C ps -C java, 通过进程名查找进程
	+ -a 显示终端上的所有进程
	+ -r 显示正在运行的进程
	+ -u 以用户为主的格式显示
	+ -x 显示所有程序，不区分终端
	+ ps -ef
		+ UID PID PPID C(cpu) STIME TTY TIME(占用cpu的时间) CMD
	+ ps -lf pid  （显示线程）


+ free 查看内存使用情况
	+ free -m (以mb为单位显示)
	+ free -h ()


+ lsof
	+ lsof -p pid 进程打开的文件
	+ lsof -i:8080 打开某端口的进程
+ 查某个进程打开的端口号
	+ lsof -p pid|grep LISTEN
	+ netstat -nltp |grep pid

+ kill -l 显示可用的信号 (信号感觉都是与进程关闭有关系的)
	+ kill -9 pid 	发送编号为9的信号，强制杀死进程 SIGKILL 
	+ kill -2  SIGINT  ctrl+c
	+ kill -15 (kill的默认选项) SIGTREM
+ SIGHUP：  终端挂起或者控制进程终止，通知session内的进程不再与终端关联
+ SIGINT: 进程中断， ctrl+C可以产生
+ SIGQUIT： 进程退出，ctrl+\  产生coredump文件
+ SIGILL:  进程在执行非法指令，比如尝试执行数据段
+ SIGSEGV: 段错误， 无效的内存引用
+ SIGTRAP： 由断点指令或其它陷阱（trap）指令产生. 由debugger使用
+ SIGKILL: 杀死进程， kill -9 pid， 不能忽略 uncatchable
+ SIGTERM: 杀死进程，可以被进程处理（忽略）, kill pid
+ SIGSTOP: 暂停进程， 不能忽略 uncatchable
+ SIGCHLD: 当子进程退出时，发送给父进程的信号。父进程可以waitpid处理，或者注册信号处理函数处理。父进程不处理，子进程变为僵尸进程。
+ SIGALRM:  timer time out
+ SIGABRT： 丢弃执行进程  调用abort函数生成的信号
+ SIGPIPE: 管道破裂: 写一个没有读端口的管道
+ SIGTSTP: 终止进程， ctrl+z
+ 让进程在后台可靠的运行
  + 用户注销或者网络中断时，终端会收到HUP信号从而关闭所有子进程。
  + 让进程忽略HUP信号 nohup
  + 让进程在新的会话中，从而不成为此终端的子进程 setsid, 让运行的程序在新的session中
+ 用户自定义信号 SIGUSR1  SIGUSR2 redis中aof子进程 与 父进程通讯使用


task_struct结构体对应与进程相关的信息

+ 创建进程 fork 子进程从父进程上继承了哪些内容？子进程与父进程的区别与联系
+ fork() pthread_create() 最后在linux中都是调用do_fork， copy_process方法
	+ linux中创建线程与进程均需要走到copy——process方法
	+ 进程与线程，对于系统来说都是task， 与task_struct对应
 	+ do_fork
		+ copy_process
		+ 复制的内容与参数flag相关
  


---

+ 软连接，与硬连接有什么区别
	+ ln -s source dist # 软连接 文件用户数据块中存放的内容是另一文件的路径名的指向，则该文件就是软连接
	+ ln source dist # 硬连接 硬链接是有着相同 inode 号仅文件名不同的文件
+ 文件删除的条件
  + 文件有2个计数器，i_count（占有的进程数）, i_nlink（硬链接数），同时为0时删除文件

+ proc目录下 查看进程的虚拟地址空间布局  cat /proc/pid/maps vm_area_struct
+ numa 架构（cpu 内存 非均匀访问） 多 CPU 系统下共享 BUS 带来的性能问题

elf格式的文件
+ core dump file
+ 可执行文件
+ 共享文件
+ 重定向文件

size 命令 输出可执行文件中的代码段，数据段，bss段的大小， 不包含堆和栈

.bss段不存在程序文件中，仅存在与内存中

cpu 寻址方式
寻址： 指令中操作数所在的真实位置，cpu怎么去找指令中的数据

七种
+ 立即数寻址（不用寻址）
+ 寄存器寻址 （操作数存在了寄存器里，不用到内存寻址）
+ 直接寻址 （直接给出操作数的内存地址）
+ 寄存器间接寻址 （操作数的地址在寄存器，操作数在内存中）
+ 寄存器相对寻址 （寄存器的内容+给定的位移量之和等于操作数的内存地址）
+ 基址加变址寻址 （基址寄存器+变址寄存器内存之和等于操作数的内存地址）
+ 相对基址变址寻址 （基址寄存器+变址寄存器+给定的位移量等于操作数的内存地址）

---

+ 计算机运算小数运算不准确的原因
   + 因为一些十进制的小数无法转换为二进制数
   + 循环小数 1/3 = 0.33333
+ 指针的值表示内存起始地址， 指针的类型表示从该起始地址到终止地址的内存尺寸

---



OS 提供了多用户多进程的良好稳定环境，但有一些技术想pass os...
+ kernel bypass 绕过内核  intel 开发的DPDK SPDK
	+ 解决内核态上的网络栈上性能问题的一种思路
	+ 避免内核：包拷贝，系统调用，线程调度等性能损耗
	+ dpdk ring buffer 
		+ array (内存连续存储，对cpu cache line友好)
		+ size power of 2 for convinient calculate
		+ padding struct obj for cache line 对齐
		+ CAS (ABA 问题 额外增加stamp，类似version判断是否出现ABA)