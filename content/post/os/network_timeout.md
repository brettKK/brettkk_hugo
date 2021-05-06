---
title: "加密相关"
date: 2021-05-05T11:33:56+08:00
draft: false
isCJKLanguage: true
markup: mmark
tags: ["encrypt"]
series: [""]
categories: ["技术"]
---

+ 散列
+ 对称加密
+ 非对称加密
+ SSL / TLS
+ 密钥的管理

---

### 散列
有md5, sha256方式， 主要保证2点： 数据没被串改；用户密码散列存储 不可逆。

### 对称加密
有des, aes 方式， 计算性能好，但密钥的安全依赖于双方，不安全。

### RSA

1977年由MIT三位学者发明。

相关概念
+ 互质 两个数的最大公约数是1。
$$gcd(a,b)=1$$
两数互质与两数是不是质数没有关系。

+ p与q互质，剩余定理得出
$$\Phi \left( pq\right) =\Phi(p)*\Phi(q)$$

+ 欧拉函数 $$\Phi\left( n\right)$$ 在小于等于n的正整数之中，与n构成互质关系的数的个数$$\Phi(n)$$

n是质数时， $$\Phi(n) = n - 1$$

n不是质数时，n分解为质数相乘 $$n = a*b*c..$$， $$\Phi \left( n\right)=\Phi \left( a\right) *\Phi \left( b\right)....$$

$$\Phi \left( p^{k}\right)=p^{k}-p^{k-1}＝p^{k}(1-\dfrac {1}{p} ) $$
$$\Phi \left(n\right)=n*(1-\dfrac {1}{p_{1}} )*(1-\dfrac {1}{p_{2}})...(1-\dfrac {1}{p_{m}})$$

+ 


#### X.509标准

公钥证书的标准，定义证书的格式
CA 

### SSL /TLS

https的原理， 握手阶段
+ 

传输阶段时使用对称加密进行通信

网络超时的类型

+ zero copy
  + tradeoff between performance and security technologies.
  + implementation:  1 memory remapping, 2 shared buffers, 3 different system calls, and hardware support.
  + https://pdfs.semanticscholar.org/6a35/60046cb8d3258669c86072a7cab05e1d2300.pd