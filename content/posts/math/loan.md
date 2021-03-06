---
title: "等额本息"
date: 2021-04-05T11:33:56+08:00
draft: false
isCJKLanguage: true
tags: ["math"]
series: [""]
categories: ["life"]
---


## 涉及的变量
+ 本金M
+ 期数N
+ 利率R
+ 等额本息X

---

## 公式推导

当N=1时，  $$ M*(1+R) = X $$ ,    $$ M_{本1} = \frac{X}{1 + R} $$

当N=2时， $$ (M - \frac{X}{1 + R})*(1+R)^2 = X $$ 

调整： $$ M = \frac{X}{1+R} + \frac{X}{(1+R)^2} $$

...

进而： $$ M = M_{本1} + M_{本2} + ... ++ M_{本n} $$

则： $$ M = \frac{X}{1+R} + \frac{X}{(1+R)^2}  + ... + \frac{X}{(1+R)^n}  $$

构造： $$ M*(1+R) = X + \frac{X}{1+R} + ... + \frac{X}{(1+R)^(n-1)} $$

错位相减： $$ M * R = X * (1 - \frac{1}{(1+R)^n}) $$

<br/>

最后本月应还的等额本息的数量为： $$ X = \frac{MR(1+R)^n}{(1+R)^n - 1}$$
