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


+ 欧拉定理 

两个正整数a与n互质， 则：

$$a^{\Phi \left(n\right)}\equiv 1\left( mod\ n\right)$$

+ 费马小定理
当n为质数时，则：$$a^{n-1}\equiv 1\left( mod\ n\right)$$

+ rsa 生成过程
  + 找两个大质数p, q
  + 求乘积 n=p * q， $$\Phi(n) = \Phi (p*q)=(p-1)*(q-1)$$
  + 在0 ～ $$\Phi(n)$$ 中，找一个e，与$$\Phi(n)$$互质
  + 找一个d， 使得 $$e*d \equiv 1(mod\ \Phi(n))
  + $$ (n, e)$$ 公钥， $$(n, d)$$ 私钥

+ 加密公式 $$x^e mod n = y $$
+ 解密公式 $$y^d mod n = x $$

+ rsa 证明

$$
y=x^e - kn \tag{1}\\
$$

$$
y^d \equiv x (mod\ n) \tag{2}\\
$$

$$
(x^e - kn)^d \equiv x (mod\ n) \tag{3}\\
$$

左边多项式展开：
$$
x^ed-m_{1}x^{e(d-1)}kn+m_{2}x^{e(d-2)}(kn)^2...m_{n}(kn)^{d} \tag{4}\\
$$

$$x^{ed}$$ 不含n， 所以只需要证明：
$$
x^{ed} \equiv x (mod\ n) \tag{5}\\
$$

而ed 等于：
$$
ed \equiv 1 (mod\ {\Phi (n)}) \tag{6}\\
$$

$$
ed=h{\Phi(n)})+1 \tag{7}\\
$$

$$
x^{h{\Phi \left(n\right)}}*x\equiv x (mod\ n) \tag{8}\\
$$

若x与n互质，根据欧拉定理 得证
$$
x^{h{\Phi \left(n\right)}}\equiv 1 (mod\ n)  \tag{9}\\
$$

$$
x^{h{\Phi \left(n\right)}}x\equiv x (mod\ n) \tag{10}\\
$$

若x与n不互质时， 稍微复杂 todo


----

+ 同样的明文，同样的公钥，每次加密的结果不会一样。11字节都为随机数 填充使用
  + 填充协议PKCS #1 v1.5: "填充后数据" = "00" + "数据块类型" + "填充字符串" + "00" + "原始数据"
    + BT (数据块类型) 00: 填充字符串全为00；01：全为FF； 02:针对公钥时设置BT=02，伪随机字符串，这能解释上面公钥加密的结果会不一样；

#### X.509标准

公钥证书的标准，定义证书的格式，包含CA机构，公钥等。

### SSL /TLS

https的原理， 握手阶段
+ 客户端给server发送支持的对称加密算法
+ server选择支持的对称加密算法和版本 （确定hash算法，对称加密算法）
+ server发送证书（包含公钥）
+ client验证证书后， 用rsa的公钥发送对称加密的密钥给server
+ server返回ack，握手完成

后续的传输阶段时使用对称加密进行通信

### 密钥的管理

+ 加解密运算 与 存储分离
+ 访问鉴权
+ 数据监控
+ root key


---

### 同态加密（Homomorphic Encryption）



全同态 支持密文上任意次，任意类型的计算

层次同态 支持有限次数的加法和乘法

部分同态 只支持加法或者乘法

可算不可见

在密文空间进行运算后与明文空间上仍然一对一，能解密。

只有不对原文进行填充的原始RSA才满足乘法同态性质。

rsa encrypt $$ c= Enc(m) = m^e mod n $$

rsa decryp $$ Dec(c) = c^d mod n $$

$$Enc(m_1) * Enc(m_2) = (m_1 * m_2)^e = Enc(m_1 * m_2)$$