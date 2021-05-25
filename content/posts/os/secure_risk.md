---
title: "安全与风控"
date: 2021-05-05T11:33:56+08:00
draft: false
isCJKLanguage: true
markup: mmark
tags: ["encrypt"]
series: [""]
categories: ["技术"]
---



---

### 漏洞

起因是信任了用户的任意输入。

#### sql注入

危害： 执行任意sql，可能出现数据泄漏风险； 避免：使用ORM框架推荐的sql预编译的方法执行sql。
即使使用了ORM框架的sql预编译的方法，框架本身有支持执行raw sql的方法。但还在一些order by, group by的参数由用户控制时会出现sql注入漏洞。
进一步，增加白名单机制，对于排序和分组场景下，可以对传入参数预设白名单。

避免： 禁止拼接sql并执行，用官方推荐用法，不信任用户的输入。

#### SSRF漏洞

RCE 任意代码执行

XSS漏洞 前端漏洞执行恶性js代码向后端

---