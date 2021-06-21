---
title: "加密相关"
date: 2021-05-20T11:33:56+08:00
draft: false
isCJKLanguage: true
tags: ["encrypt"]
math: true
series: [""]
categories: ["技术"]
---
## 对称加密

有des, aes 方式， 计算性能好，但密钥的安全依赖于双方，密钥有配送安全问题。

散列: func后不可逆， 不是加密方式。有md5, sha256。保证数据没被串改；用户密码散列存储。

---
## RSA

### RSA背景
1977年由MIT三位学者发明。

### 相关数论概念

+ 互质 两个数的最大公约数是1。 $$gcd(a,b)=1$$ 两数互质与两数是不是质数没有关系。

+ p与q互质，剩余定理得出 $$\Phi \left( pq\right) =\Phi(p)*\Phi(q)$$

+ 欧拉函数 $$\Phi(n)$$

在小于等于n的正整数之中，与n构成互质关系的数的个数

n是质数时， $$\Phi(n) = n - 1$$

n不是质数时，n分解为质数a,b,c...相乘 $$n = a * b * c\cdots$$

$$\Phi(n)=\Phi( a) *\Phi( b)...$$


+ 欧拉定理 

两个正整数a与n互质， 则：

$$a^{\Phi \left(n\right)}\equiv 1\left( mod\ n\right)$$

+ 费马小定理
  
当n为质数时，则：$$a^{n-1}\equiv 1\left( mod\ n\right)$$

### rsa 生成过程

找两个大质数p, q
 
求乘积 n=p * q， $$ \Phi(n) = \Phi (p * q)=(p-1)*(q-1) $$

在0 ～ n的欧拉函数之间找一个e，使得e与欧拉函数互质

找一个d， 使得 $$e * d \equiv 1 mod\Phi(n) $$

则(n, e)公钥， (n, d)私钥

+ 加密公式 $$x^e mod n = y $$
+ 解密公式 $$y^d mod n = x $$

### rsa 证明

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


### 工程实现

+ 明文需要分段加密，原因是明文大小不能超过n的限制。
+ 同样的明文，同样的公钥，每次加密的结果不会一样。因为随机数填充。
  + 填充协议PKCS #1 v1.5: "填充后数据" = "00" + "数据块类型" + "填充字符串" + "00" + "原始数据"
  + BT (数据块类型) 00: 填充字符串全为00；01：全为FF； 02:针对公钥时设置BT=02，伪随机字符串，这能解释上面公钥加密的结果会不一样；

---

## TLS

https的原理， 握手阶段
+ 客户端给server发送支持的对称加密算法
+ server选择支持的对称加密算法和版本 （确定hash算法，对称加密算法）
+ server发送证书（包含公钥， X.509标准公钥证书的标准，定义证书的格式，包含CA机构，公钥等）
+ client验证证书后， 用rsa的公钥发送对称加密的密钥给server
+ server返回ack，握手完成

后续的传输阶段时使用对称加密进行通信


### TLS1.2 and TLS1.3

Transport Layer Security 1.3 参考链接 https://www.rfc-editor.org/rfc/pdfrfc/rfc8446.txt.pdf 

#### TLS 背景

提供端到端之间提供安全通信通道。 

依赖： 依赖下层协议的可靠性和顺序数据流。

#### TLS 特性

认证Authentication方式： RSA, the Elliptic Curve Digital Signature Algorithm (ECDSA), symmetric pre-shared key (PSK)。

机密

完整性

### TLS two primary components

+ 握手协议

协商密码模式和参数，并建立共享密钥材料

+ 记录协议

### TLS1.2 与 TLS1.3 的区别

TLS1.2  -- performance and security --> TLS 1.3

+ 增加O-RTT

O-RTT for repeat clients，1-RTT for native clients，TLS 1.3 shorter handshake。

+ static rsa, dh被移除， 增加的密钥交互机制均保证前向数据安全。


### DH &&  ECDH

DH (Diffie-Hellman) : 能在非安全的信道中安全的交换密钥，用于加密后续的通信消息。

### DH 交互流程

```
A: 构建密钥对：private key1 和 public key1
A->> B: 公布自己的公钥: public key1
B: 使用public key1构建自己的密钥对 private key2 和 public key2；
B-->> A: 返回自己的public key2；
A: 使用private key1 和 public key2 构建本地密钥；
B : 使用private key2 和 public key1构建本地密钥；
```

#### DH 形式话描述

算法形式化描述： 已知变量对(p, q), 

alice 本地生成随机数a， 计算$$ G_a= (p^a) mod q$$

bob 本地生成随机数b， 计算$$ G_b = (p^b) mod q $$

进而 alice 计算出通信密钥为： $$ S_a = G_b^a mod q $$  

bob 也计算出通信密钥为：   $$ S_b =G_a^b mod q $$

需证明 $$S_a == S_b$$



#### ECDH

ECC

给定椭圆曲线上的一个点P，一个整数k，求解Q=kP很容易；给定一个点P、Q，知道Q=kP，求整数k确是一个难题。ECDH即建立在此数学难题之上。

```
Alice生成随机整数a，计算A=a*G。 #生成Alice公钥

Bob生成随机整数b，计算B=b*G。 #生产Bob公钥
 
Alice将A传递给Bob。 Bob将B传递给Alice

进而：
Alice计算  Q = a * B  , Bob就算 Q = b * A

证明： Q=b * A = b * (a * G) = a * b * G = a * (b * G) = a * B 
则： b * A = a * B , 双方密钥相同。
```

tls 1.3 中使用 x25519椭圆曲线 $$ B * y^2 = x^3 + A * x^2 + x $$



## 密钥的管理

kms系统： 信封加密， 密钥轮替

+ 加解密运算 与 存储分离
+ 访问认证
+ 数据监控
+ root key

---

##   同态加密（Homomorphic Encryption）

全同态 支持密文上任意次，任意类型的计算

层次同态 支持有限次数的加法和乘法

部分同态 只支持加法或者乘法

可算不可见

在密文空间进行运算后与明文空间上仍然一对一，能解密。

只有不对原文进行填充的原始RSA才满足乘法同态性质。

rsa encrypt $$ c= Enc(m) = m^e mod n $$

rsa decryp $$ Dec(c) = c^d mod n $$

$$Enc(m_1) * Enc(m_2) = (m_1 * m_2)^e = Enc(m_1 * m_2)$$


## ABE 

下一代公钥密钥的先进机制
+ 一对多的加解密方式
+ 访问和数据加密融为一体

Key-Policy ABE： 策略嵌入用户密钥中，属性集合嵌入于密文中。

CP-ABE： 策略嵌入密文，通过设定策略去决定拥有哪些属性的人能够访问这份密文

公有云上的数据加密存储与细粒度共享。

联邦学习：为解决AI技术落地的数据安全隐私问题。