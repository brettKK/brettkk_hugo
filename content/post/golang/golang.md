---
title: "Golang"
date: 2021-05-04T18:54:41+08:00
draft: false
isCJKLanguage: true
---

资料
+ https://github.com/gocn/knowledge
+ awesome-go


+ 数组的初始化 （元素类型， 数组大小）
    + cmd/compile/internal/types.NewArray
    + cmd/compile/internal/types.NewSlice
    + 2种初始化方式
        + 显式的指定数组的大小 // arr := [3]int{1, 2, 3}
        + [...]T 声明数组 // arr := [...]int{1, 2, 3}

+ goroot/src/syscall --》 建议使用 golang.org/x/sys 


