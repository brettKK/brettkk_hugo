---
title: "linux-tcp"
date: 2021-05-05T11:33:56+08:00
draft: false
isCJKLanguage: true
markup: mmark
tags: ["linux", "protocol"]
series: [""]
categories: ["技术"]
---




tcp报文结构图
<img src="/os/tcp-header.png" width = "400" /><br>

+ URG: 表示紧急指针（urgent pointer）是否有效
+ PSH: 提示接收端应该立即从TCP接收缓冲区中读走数据，为接收后续数据腾出空间
+ RST：复位报文段，用来关闭异常的TCP连接 （重要） （4种情况会出现： 端口未打开，请求超时，提前关闭，已经的socket上收到数据）
  + A想B发起tcp连接，但是B上未监听相应的端口，此时B的os的网络模块会发送RST包给A端
  + syn攻击时防火墙设置，向服务器发送RST包，清除未连接队列的数据包，保证避免server出现拒绝服务的情况
  + RST攻击场景，攻击则伪装为客户端向server发送RST包，导致连接被强制关闭
+ window：是TCP流量控制的一个手段， 告诉对方本端的TCP接收缓冲区还能容纳多少字节的数据，这样对方就可以控制发送数据的速度(sender在未收到ack时最大发送的字节数)
+ checkSum: 校验数据包是否损坏

两个线程如何同时监听一个端口,SO_REUSEPORT 参数

+ 4次挥手： A 和 B 打电话，通话即将结束后，A 说“我没啥要说的了”，B回答“我知道了”，但是 B 可能还会有要说的话，A 不能要求 B 跟着自己的节奏结束通话，于是 B 可能又巴拉巴拉说了一通，最后 B 说“我说完了”，A 回答“知道了”，这样通话才算结束
 
tcp握手与挥手过程图（基本图，实际中可以不是这样，例如关闭时双方同时发起关闭，双方都会到达time_wait状态）

<img src="/os/tcp-connect.png" width = "600" /><br>
<img src="/os/tcp-close.png" width = "600" /><br>

⚠️
wireshake上的seq和ack number是relative number，0和1是为了方便查看，实际情况是其他数值

TIME-WAIT状态的理解：
+ 持续时间未2MSL，一个数据包在网络中的最长生存时间为MSL(30S)
+ 客户端回复的ACK丢失，服务器端会在超时时间到来时，重传最后一个FIN包
+ ACK和FIN在网络中的最长生存时间就为2MSL，这样就可以可靠的断开TCP的双向连接

服务器解决TIME——WAIT（占用端口过多，导致无可用的端口）的方式3种：
+ 保证客户端主动关闭
+ 关闭时使用RST方式
+ 对于处于TIME-WAIT状态的tcp允许重用，修改内核参数(java服务器访问mysql集群，主动关闭，time-wait状态在服务器，他是发起方，可以复用tcp连接)
+ 当利用http发送请求时，keep-alive减少tcp连接数，进而减少time-wait状态的出现

---
TCP如何实现可靠传输（数据应答机制+数据超时重传）
+ 数据包都序列号，保证接收方能恢复数据包的顺序，去重等
+ 超时重传（确认和重传）， 也有快速重传，不用等待超时时间
  + 超时时间大于RTT（一次往返时间）
  + 每个被发送的数据包均有计时器，只有超过计时器的重传时间后，才会重发数据包
  + 1988 最早的超时重传RTO = min(ubound, max(lbound, beta * SRTT)); SRTT= (alph * SRTT) + (1- alpha) * RTT ; alpha=0.8, 0.9 beta=1.2, 1.3
  + RTO 比 RTT 大很多时，若出现网络丢包，重传速度很慢； RTO假设比RTT小时，会出现大量的重传； 所以RTO的时间比较考量

TCP如何实现流量控制
+ 当接受方处理不过来时，接受方缩小窗口，告知发送方
+ 以字节为单位的滑动窗口技术
  + recv的窗口大小决定sender的窗口大小（窗口以字节为单位）
  + sender缓存中的数据包必须收到recv的确认数据包后才能删除
+ 当接收方来不及处理发送方的数据，能提示发送方降低发送的速率，防止包丢失

TCP如何实现拥塞控制
+ 当网络拥塞时，动态改变发送窗口大小，减少数据的发送

---

基于TCP的攻击
+ SYN泛洪攻击
  + 发送大量的SYN包，导致server上未连接队列满，导致拒绝正常的连接服务
  + 解决方法：
    + SYN网关， 防火墙收到SYN包转给server，防火墙收到server的ack/syn后，转给client；同时防火墙以client的名义直接给server发送ack，完成3次握手。 server上连接数远大于未连接数。
    + 被动式SYN网关，防火墙设置小的超时时间，超时时向server发送RST包，导致关闭半连接
    + SYN中继， 完全有防火墙做代理，防火墙与client握手完成表明正常后，在于server完成3次握手
+ RST攻击 （ip头+tcp头 40字节）
  + 场景： A--》B， C（伪装A）--》B， 发送RST包，或者正在连接时发送SYN包，B会清空A请求过来的数据包，强制关闭连接
  + C如何伪装A （tcp四要素：源ip（容易知道），源port，目的ip（容易知道），目的port（容易知道））
    + 关键是知道源port（发现os分配时的规律）和序列号（需要在B的滑动窗口内，不然会被丢弃）
    + 大量发送数据包（带宽不是问题，RST包很小）保证一个落在B的滑动窗口内 ，sequence number为32位，滑动窗口为65535，相除为65537个数据包就有一个落在区间内
---
