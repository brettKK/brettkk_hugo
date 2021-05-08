---
title: "linux-shell"
date: 2021-05-05T11:33:56+08:00
draft: false
isCJKLanguage: true
markup: mmark
tags: ["linux"]
series: [""]
categories: ["技术"]
---

+ 单引号，双引号，反引号区别
	+ 单引号之间的内容原封不动的输出
	+ 双引号取消空格的作用， 特殊符号含义会保留
	+ 反引号，用来包装命令的结果，多用$()
```
	将一个命令的执行结果赋给变量
	A=`date` 将date的执行结果赋值给变量A
	B=$(ls -l) 将ls -l这个命令给变量B
```
+ shell 数组
	+ 定义数组	array_name=(v0 v1 v2 v3)
	+ 读取数组 value1=${array_name[1]}
	+ 获取所有元素	@，*
	+ 获取数组长度 echo ${#array_name[ * ]}

+ shell 传递参数
	+ $n 位置变量
		+ n为一个数字，0为执行的文件，1位第一个参数，类推
	+ 特殊字符
		+ $# 传递到脚本的参数个数
		+ $* 以字符串的显示参数 “$1..$n”
		+ $@ 与$* 类似 加引号
		+ $$ 这个程序的pid
		+ $! 执行上一个后台程序的pid
		+ $? 执行上一个指令的返回值

### awk

awk 脚本基本结构 awk 'BEGIN{ print "start" } pattern{ commands } END{ print "end" }' file  

使用AWK做文本处理， 语句由模式和动作组成
+ 模式，决定动作的触发条件和时间
+ 动作
  + BEGIN， 设置计数和头部信息， 在任何动作之前进行
  + END， 输出统计结果， 在完成动作之后执行

默认分隔符为空格， 也可以-F改变为其他符号，例如冒号，-F：

例自定义日期显示
date|awk '{print "year:"$1 "month:"$2}'

自定义输出一些列
awk 'BEGIN {print "begin"} {print $1, $2, $3} END{print "end"}' file

满足某个模式的行
awk '$2>9 {print $0}' file

awk '{if($1=="kk" && $2>5) print $0}' file

---

### sed

替换s命令
sed 's/old/new/' file

-n -p 一起使用只打印发生替换的行
sed -n 's/old/new/p' file
显示第三行
sed -n '3p' file

全部替换g命令
sed 's/old/new/g' file

删除d命令
sed '2d' file

插入i命令
追加a命令

---

+ /etc/profile
	+ 为系统中每个用户设置环境信息，只在用户第一次登陆时，被执行
	+ 并从/etc/profile.d目录中搜索shell的配置
+ /etc/bashrc
	+ bash shell被打开，就被执行
+ ～/.bashrc
	+ 每次打开shell时，被执行
+ ~/.bash_profile
	+ 用户登陆时仅执行一次，执行用户的bashrc

+ 重定向
	+ runApp < data.in > results.out
	+ > /dev/null 丢弃输出信息
	+ 2>&1 错误输出重定向到标准输出中
	+ >> a.log  追加写文件
	+ &>a.log  将标准输出和错误输出的组合重定向到a.log里
	+ 单独的 & 将任务发送到shell后台执行
		+ jobs 列出在后台执行的所有任务
		+ fg 任务编号 将任务带回前台
		+ ctrl+z 暂停进程， bg 将任务继续带回后台执行


+ 查找某个字符串
  + 命令模式:／strings ，光标位置向下查找
  + :?strings, 光标位置向上查找
+ 在整个文件中替换特定字符串
  + :%s/old_word/new_word/g // %s/old_word/new_word/gc  每次替换会确认一下
  + g/old_word/s//new_word/g



+ zero copy
  + tradeoff between performance and security technologies.
  + implementation:  1 memory remapping, 2 shared buffers, 3 different system calls, and hardware support.
  + https://pdfs.semanticscholar.org/6a35/60046cb8d3258669c86072a7cab05e1d2300.pd