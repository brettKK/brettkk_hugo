---
title: "golang"
date: 2021-05-04T18:54:41+08:00
draft: false
isCJKLanguage: true
tags: ["golang"]
series: [""]
categories: ["技术"]
---

### 资料
+ https://github.com/gocn/knowledge
+ awesome-go

### 模糊点
+ 数组的初始化 （元素类型， 数组大小）
    + cmd/compile/internal/types.NewArray
    + cmd/compile/internal/types.NewSlice
    + 2种初始化方式
        + 显式的指定数组的大小 // arr := [3]int{1, 2, 3}
        + [...]T 声明数组 // arr := [...]int{1, 2, 3}

+ 函数接受者为指针类型
  + 函数可以修改接受者的内容
  + 避免方法调用时变量的拷贝，而是指针的拷贝
+ go严格上只有复制拷贝传递
  + 总是创建副本按值传递，只不过这个副本可以是变量的副本，也可以是指针的副本

+ unsafe
+ unsafe.Pointer()
  + T1指针与 T2指针类型之间的转换
  + 修改内存数据 uintptr 常用于与 unsafe.Pointer配合，用于做指针运算

+ 反射 Reflection (interface对象=（Type, Value）反射类型 reflect.Type, reflect.Value)  三点
    + 反射可以将interface类型的变量 转换为 反射对象（reflect.Type reflect.Value）
    + 反射可以将反射对象 还原为interface对象
    + 反射对象可以修改，提前是value值是变量的指针  var x float= 3.1 ; v:= reflect.ValueOf(&x); v.Elem().SetFloat(3.4);
+ 动态调用方法， reflect.Value(interface) -> reflect.Value->value.MethodByName->Call

+ golang 内存管理
    + 内存分配原理
        + src/runtime/mheap.go:mspan
    + 内存回收原理
        + 三色标记法 （未使用，已使用，带处理）
    + 逃逸分析（对象是否被函数外面引用 例如闭包； 对象过大也可能在堆中）
        + 逃逸分析的目的是决定变量的内存分配地址是在栈还是堆
        + 逃逸分析在编译阶段完成的
        + 传递指针真的比传值高效吗？不一定 由于指针传递会产生逃逸，可能会使用堆，增加GC负担。
        + go build -gcflags=-m // escapes to heap


+ nil
    + what is nil
        + nil is a kind of zero
    + what is nil in go
        + 各种类型的零值
        + zero value of all type value
            + bool false, number 0, string ""
            + pointer, slice, map, channel, function, interface{} ==> nil
            + structure{}
        + kind of nil 
            + interface{} = (type, value) ==> (nil, nil) == nil 
    + what does nil mean
    + is nil useful ? yes

+ avo
    + github.com/mmcloughlin/avo

### go 汇编
+ 汇编风格分为2类
  + intel风格
  + AT&T风格(左 --> 右) go是plan9， 属于此类

+ 助记符    call, add
+ register  ax, sb, pc
+ 立即数 $1 , $0x100
+ 内存寻址 （R1）

+ plan 9
    + https://lrita.github.io/2017/12/12/golang-asm/#go%E5%87%BD%E6%95%B0%E8%B0%83%E7%94%A8
    + https://lrita.github.io/images/posts/go/GoFunctionsInAssembly.pdf
    + https://github.com/buptbill220/gotls/tree/master/interesting
    + 查看当前cpu架构的指令集合 cat /proc/cpuinfo | grep flags | head -1
    + 指令查询： http://68k.hax.com/
    + 命令查询：https://9p.io/magic/man2html/1/8a
    + plan9 文档： https://9p.io/sys/doc/
+ https://defuse.ca/online-x86-assembler.htm#disassembly2
  +  x86 or x64 assembly instructions encoding<==> decoding binary code
  +  It uses GCC and objdump behind the scenes

+ goroot/src/syscall --》 建议使用 golang.org/x/sys 


