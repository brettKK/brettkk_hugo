---
title: "http, curl, dns"
date: 2021-05-05T11:33:56+08:00
draft: false
isCJKLanguage: true
markup: mmark
tags: ["linux", "protocol"]
series: [""]
categories: ["技术"]
---
### http 

+ http header 常用字段
  + content-type (如何解析body)
    + application/x-www-form-urlencoded
    + multipart/from-data, 上传文件
    + application/json, text/plain
    + text/xml, XML格式 （post请求的4种格式）
  + connection: keep-alive, 持久连接（相对于用完就关闭的短连接而言）
  + keep-alive-time: 300, 指连接维持的时间
  + content-length: 120, 指明响应的内容有120字节 (出现等于-1，但是body有数据的情况？)
  + transfer-encoding: chunked， 响应内容采用分块传输，最后一个分块为0表示传输结束
+ http keep-alive 与 tcp keep alive 区别与联系
  + http 用于连接复用
  + tcp 用于保证连接存活，发送探测包确认连接是否存活

+ 有时ping url结果为unkown host，但是curl url可以通
  + 有可能对方防火墙屏蔽ping的ICMP协议

---
### curl命令

+ 不带参数 curl  xx.com 发get请求
+ curl -X POST xx.com 发post请求
  
+ -I only response header
  +  curl -I www.google.com
  + (curl -I 获取http响应的头部)
+ -v 显示http通信过程 --trace
  + curl -v http://www.example.com/index.html
  + > 发送的请求， <接收的响应
  + RFC2616

+ curl -u username :password URL
+ curl 查询单词
	+ curl dict://dict.org/d:xxx 查询xxx的含义
	+ curl dict://dict.org/show:db 列出可用的词典
	+ curl dict://dict.org/d:xx:foldoc 利用foldoc查询xx的含义
+ 给curl 设置代理
	+ curl -x proxyserver.test.com:9090 google.com
+ 将cookie保存到本地
	+ curl -D localCookie google.com
	+ curl -b localCookie google.com 使用上次的cookie
	+ curl -c cookie url 保存cookie到文件
+ 默认curl使用get方式 发送表单
	+ curl -u username github.com
+ 可以通过--data/-d方式使用post
	+ curl -u username --data "param1=value1&param2=value2" github.com
	+ curl --data @filename github.com 将文件中的数据传给服务器
+ User Agent (-A , -H)
	+ curl --user-agent "user-agent" url
+ cookie (-b)
	+ curl --cookie "name=xx" url
	+ cookie值可以从response header的·Set-Cookie·字段得到
+ 自行加request header信息
	+ curl --header "" url
+ curl -d '{"key1":"value1", "data":{"data1":"value1", "data2":"value2"}}' 'https://xxx.com/a/b'
+ curl -F 'business_type=one' -F 'data=@/Users/kk/a.csv'  x.x.x.x:xx/a/b curl发送文件
+ curl -G -d 'q=xx' xx.com -d ''//xx.com?q=xx

+ ping icmp协议 封装在IP包里
	+ 类型0 8 请求与响应
	+ 3是不可达 后面有代码号 
  	+ 5种情况 
    	+  ip中的网络号-网络不可达 
    	+  ip中的主机号-主机不可达 
    	+  防火墙不允许tcp连接-协议不可达
    	+  端口没有进程监听-端口不可达


---
### dig 

+ dig查询DNS相关信息的工具
  + 显示internet上13个根域服务器
  + dig @dnsserver name querytype
    + 用google的dns查询baidu的A记录
    + dig @8.8.8.8 www.baidu.com A
+ DNS记录的类型
  + A记录       指定域名对应的IP地址
  + AAAA记录    指定域名对应的IPV6地址
  + CNAME记录   别名解析 域名1（CDN加速域名） -》域名2 -》IP
  + NS记录      域名服务器记录  将域名解析交给其他DNS服务器



---

+ 访问github.com时与DNS解析有关的一个问题

问题：fatal: unable to access 'https://github.com/brettKK/shell.git/': Failed to connect to github.com port 443: Operation timed out

方法：vim /etc/hosts ==> find github.com to delete